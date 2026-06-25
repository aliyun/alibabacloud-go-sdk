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
  // The job template ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ee1a7a07-abcb-4652-a1d3-2d57f415****
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // The startup command, which must be an executable that exists in the container. Example:
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
  // Based on this example, `Command` is `echo` and `CommandArgs` is `["abc", ">", "file0"]`.
  // 
  // example:
  // 
  // echo
  Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
  // The arguments for the **Command*	- parameter. The value must be a string that represents a JSON array. Format:
  // 
  // `["a","b"]`
  // 
  // In the preceding example for the `Command` parameter, `CommandArgs` is `["abc", ">", "file0"]`. The JSON array `["abc", ">", "file0"]` must be converted to a string. This parameter is optional.
  // 
  // example:
  // 
  // ["a","b"]
  CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
  // The container environment variables. You can specify custom environment variables or reference an existing ConfigMap. For more information about creating a ConfigMap, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html).
  // 
  // - Custom configuration
  // 
  //   - **name**: The name of the environment variable.
  // 
  //   - **value**: The value of the environment variable.
  // 
  // - Reference a ConfigMap
  // 
  //   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, use the `sae-sys-configmap-all-<ConfigMap name>` format. Example: `sae-sys-configmap-all-test1`.
  // 
  //   - **valueFrom**: The source of the environment variable. Set the value to `configMapRef`.
  // 
  //   - **configMapId**: The ID of the ConfigMap.
  // 
  //   - **key**: The key that you want to reference. If you want to reference all key-value pairs, do not specify this parameter.
  // 
  // example:
  // 
  // [{"name":"envtmp","value":"0"}]
  Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
  // A customizable event ID that ensures idempotency. The system creates only one job for requests that have the same event ID.
  // 
  // example:
  // 
  // custom
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  // Arguments for starting a job deployed from a JAR package. The default startup command is:
  // 
  // `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
  // 
  // example:
  // 
  // custom-args
  JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
  // Options for starting a job deployed from a JAR package. The default startup command is:
  // 
  // `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArg`
  // 
  // example:
  // 
  // -Xms4G -Xmx4G
  JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
  // The number of concurrent instances.
  // 
  // example:
  // 
  // 3
  Replicas *string `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
  // The time to trigger the job, specified in the `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"` format.
  // 
  // example:
  // 
  // 2023-09-14T14:25:02Z
  Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
  // The startup command for a job deployed from a WAR package. Configuration is the same as for an image-based deployment. For more information, see [Configure a startup command](https://help.aliyun.com/document_detail/96677.html).
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

