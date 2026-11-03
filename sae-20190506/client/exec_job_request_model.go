// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecJobRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *ExecJobRequest
  GetAppId() *string 
  SetCommand(v string) *ExecJobRequest
  GetCommand() *string 
  SetCommandArgs(v string) *ExecJobRequest
  GetCommandArgs() *string 
  SetEnvs(v string) *ExecJobRequest
  GetEnvs() *string 
  SetEventId(v string) *ExecJobRequest
  GetEventId() *string 
  SetJarStartArgs(v string) *ExecJobRequest
  GetJarStartArgs() *string 
  SetJarStartOptions(v string) *ExecJobRequest
  GetJarStartOptions() *string 
  SetReplicas(v string) *ExecJobRequest
  GetReplicas() *string 
  SetTime(v string) *ExecJobRequest
  GetTime() *string 
  SetWarStartOptions(v string) *ExecJobRequest
  GetWarStartOptions() *string 
}

type ExecJobRequest struct {
  // The ID of the job template.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ee1a7a07-abcb-4652-a1d3-2d57f415****
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // The image startup command. The command must be an executable object that exists in the container. Example:
  // 
  // ```
  // 
  // command:
  // 
  //       - echo
  // 
  //       - abc
  // 
  //       - >
  // 
  //       - file0
  // 
  // ```
  // 
  // Based on the preceding example, Command="echo" and CommandArgs=["abc", ">", "file0"].
  // 
  // example:
  // 
  // echo
  Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
  // The arguments for the image startup command. These are the arguments required by the **Command*	- startup command. Format:
  // 
  // `["a","b"]`
  // 
  // In the preceding example, CommandArgs=["abc", ">", "file0"]. The ["abc", ">", "file0"] part must be converted to a string in JSON array format. If this parameter is not required, leave it empty.
  // 
  // example:
  // 
  // ["a","b"]
  CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
  // The environment variables for the container. You can customize environment variables or reference configuration items. To reference a configuration item, you must first create a ConfigMap instance. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
  // 
  // - Custom configuration
  // 
  //     - **name**: the name of the environment variable.
  // 
  //     - **value**: the value of the environment variable.
  // 
  // - Referenced configuration items
  // 
  //     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter sae-sys-configmap-all-<configuration item name>, for example, sae-sys-configmap-all-test1.
  // 
  //     - **valueFrom**: the reference of the environment variable. Set the value to configMapRef.
  // 
  //     - **configMapId**: the ID of the configuration item.
  // 
  //     - **key**: the key. If you reference all key-value pairs, do not set this field.
  // 
  // example:
  // 
  // [{"name":"envtmp","value":"0"}]
  Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
  // The event ID. This parameter is user-defined and used for idempotence. Only one job is created for the same event ID.
  // 
  // example:
  // 
  // custom
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  // The startup arguments for the JAR package job. Default job startup command:
  // 
  // `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
  // 
  // example:
  // 
  // custom-args
  JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
  // The startup options for the JAR package job. Default job startup command:
  // 
  // `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArg`
  // 
  // example:
  // 
  // -Xms4G -Xmx4G
  JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
  // The number of concurrent instances for the job.
  // 
  // example:
  // 
  // 3
  Replicas *string `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
  // The time when the job is triggered. Format: `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"`.
  // 
  // example:
  // 
  // 2023-09-14T14:25:02Z
  Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
  // The startup command for the WAR package deployment job. The configuration procedure is the same as that for the image deployment startup command. For more information, see [Configure a startup command](https://help.aliyun.com/document_detail/96677.html).
  // 
  // example:
  // 
  // CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
  WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
}

func (s ExecJobRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecJobRequest) GoString() string {
  return s.String()
}

func (s *ExecJobRequest) GetAppId() *string  {
  return s.AppId
}

func (s *ExecJobRequest) GetCommand() *string  {
  return s.Command
}

func (s *ExecJobRequest) GetCommandArgs() *string  {
  return s.CommandArgs
}

func (s *ExecJobRequest) GetEnvs() *string  {
  return s.Envs
}

func (s *ExecJobRequest) GetEventId() *string  {
  return s.EventId
}

func (s *ExecJobRequest) GetJarStartArgs() *string  {
  return s.JarStartArgs
}

func (s *ExecJobRequest) GetJarStartOptions() *string  {
  return s.JarStartOptions
}

func (s *ExecJobRequest) GetReplicas() *string  {
  return s.Replicas
}

func (s *ExecJobRequest) GetTime() *string  {
  return s.Time
}

func (s *ExecJobRequest) GetWarStartOptions() *string  {
  return s.WarStartOptions
}

func (s *ExecJobRequest) SetAppId(v string) *ExecJobRequest {
  s.AppId = &v
  return s
}

func (s *ExecJobRequest) SetCommand(v string) *ExecJobRequest {
  s.Command = &v
  return s
}

func (s *ExecJobRequest) SetCommandArgs(v string) *ExecJobRequest {
  s.CommandArgs = &v
  return s
}

func (s *ExecJobRequest) SetEnvs(v string) *ExecJobRequest {
  s.Envs = &v
  return s
}

func (s *ExecJobRequest) SetEventId(v string) *ExecJobRequest {
  s.EventId = &v
  return s
}

func (s *ExecJobRequest) SetJarStartArgs(v string) *ExecJobRequest {
  s.JarStartArgs = &v
  return s
}

func (s *ExecJobRequest) SetJarStartOptions(v string) *ExecJobRequest {
  s.JarStartOptions = &v
  return s
}

func (s *ExecJobRequest) SetReplicas(v string) *ExecJobRequest {
  s.Replicas = &v
  return s
}

func (s *ExecJobRequest) SetTime(v string) *ExecJobRequest {
  s.Time = &v
  return s
}

func (s *ExecJobRequest) SetWarStartOptions(v string) *ExecJobRequest {
  s.WarStartOptions = &v
  return s
}

func (s *ExecJobRequest) Validate() error {
  return dara.Validate(s)
}

