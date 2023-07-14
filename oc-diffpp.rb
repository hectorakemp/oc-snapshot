#!/usr/bin/env ruby

require 'yaml'
require 'json'

# create a diff between the last applied configuration on a cluster and the current configuration.

# oc-diff ingresscontroller -n openshift-ingress-operator
#
puts ARGV
oc_yaml = `oc get #{ARGV.join(' ')} -oyaml`

last_appl = YAML.load(oc_yaml)['metadata']['annotations']['kubectl.kubernetes.io/last-applied-configuration']
last_appl_json = JSON.parse(last_appl)

last_appl_yaml = last_appl_json.to_yaml

`echo '#{oc_yaml}' > foo.yaml`
`echo '#{last_appl_yaml}' > bar.yaml`

puts `diff foo.yaml bar.yaml`
