# Tektoncd cli TP-1.2

# Tekton Cli v0.13.0

## Features :sparkles:
* Enable auto-select support in ClusterTask/TaskRun/pipelineDescribe if only one is present 
https://github.com/tektoncd/cli/pull/1154 and https://github.com/tektoncd/cli/pull/1187
* Use --prefix-name option for tkn clustertask start https://github.com/tektoncd/cli/pull/1179
* Enable auto select support in PipelineRunDescribe if only one PipelineRun Present https://github.com/tektoncd/cli/pull/1190

## Fixes :bug:
* Modify tkn version to accept namespace as ldflag and flag https://github.com/tektoncd/cli/pull/1092
* Cancelling a pipelinerun which failed with Failed(Cancelled) throws an error https://github.com/tektoncd/cli/pull/1185
* Fix deployment fetch issue for multiple namespaces https://github.com/tektoncd/cli/pull/1186

# Tekton Cli v0.12.0

## Features :sparkles:
* Display results in the output of task/taskrun, ClusterTask and pipeline/pipelinerun describe command 
https://github.com/tektoncd/cli/pull/1076, https://github.com/tektoncd/cli/pull/1105 and https://github.com/tektoncd/cli/pull/1110
* ClusterTask start interactive mode https://github.com/tektoncd/cli/pull/1051
* Add Ability to Specify PodTemplate for TaskRun/PipelineRun https://github.com/tektoncd/cli/pull/1088
* Add --use-param-defaults option for tkn clustertask start https://github.com/tektoncd/cli/pull/1094
* Display workspaces in tkn task/clustertask/taskrun and pipeline/pipelinerun describe command 
https://github.com/tektoncd/cli/pull/1121 and https://github.com/tektoncd/cli/pull/1142
* Modify --use-param-defaults flag for pipeline start command and add e2e test https://github.com/tektoncd/cli/pull/1109

## Fixes :bug:
* Add New Line for No ClusterTasks found Output https://github.com/tektoncd/cli/pull/1101
* Add Check for Timeout Status for PipelineRun and TaskRun Cancel https://github.com/tektoncd/cli/pull/1124
* Fix the display of 0 Value for Timeout with tkn tr/pr describe https://github.com/tektoncd/cli/pull/1117
* pkg/suggestions: print help and error out https://github.com/tektoncd/cli/pull/1126
* pkg/suggestions: no help on subcommand errors https://github.com/tektoncd/cli/pull/1127
* Add EventListener Status with tkn eventlistener list https://github.com/tektoncd/cli/pull/1098
* Consistent Error Messaging for Triggers Commands https://github.com/tektoncd/cli/pull/1130
* Throw error when deleting tr/pr with non-existing task/pipeline https://github.com/tektoncd/cli/pull/1123
* Copy Over PipelineRun and TaskRun Spec for --last and use run Options https://github.com/tektoncd/cli/pull/1134
* Refactor status formatting https://github.com/tektoncd/cli/pull/1141
* Refactor taskrun description https://github.com/tektoncd/cli/pull/1144
