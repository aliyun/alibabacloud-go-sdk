// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLogRecordCustomizedHeadersEnabled(v bool) *UpdateListenerLogConfigRequest
	GetAccessLogRecordCustomizedHeadersEnabled() *bool
	SetAccessLogTracingConfig(v *UpdateListenerLogConfigRequestAccessLogTracingConfig) *UpdateListenerLogConfigRequest
	GetAccessLogTracingConfig() *UpdateListenerLogConfigRequestAccessLogTracingConfig
	SetClientToken(v string) *UpdateListenerLogConfigRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateListenerLogConfigRequest
	GetDryRun() *bool
	SetListenerId(v string) *UpdateListenerLogConfigRequest
	GetListenerId() *string
}

type UpdateListenerLogConfigRequest struct {
	// Specifies whether to record custom headers in the access log. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > You can set this parameter to **true*	- only if the access log feature is enabled by specifying **AccessLogEnabled**.
	//
	// example:
	//
	// true
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The configuration information about the Xtrace feature.
	AccessLogTracingConfig *UpdateListenerLogConfigRequestAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: (default): performs a dry run and performs the actual request. If the request passes the dry run, a **2xx HTTP*	- status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Application Load Balancer (ALB) listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s UpdateListenerLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerLogConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigRequest) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *UpdateListenerLogConfigRequest) GetAccessLogTracingConfig() *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	return s.AccessLogTracingConfig
}

func (s *UpdateListenerLogConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateListenerLogConfigRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateListenerLogConfigRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateListenerLogConfigRequest) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *UpdateListenerLogConfigRequest {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetAccessLogTracingConfig(v *UpdateListenerLogConfigRequestAccessLogTracingConfig) *UpdateListenerLogConfigRequest {
	s.AccessLogTracingConfig = v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetClientToken(v string) *UpdateListenerLogConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetDryRun(v bool) *UpdateListenerLogConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetListenerId(v string) *UpdateListenerLogConfigRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) Validate() error {
	if s.AccessLogTracingConfig != nil {
		if err := s.AccessLogTracingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateListenerLogConfigRequestAccessLogTracingConfig struct {
	// Specifies whether to enable the Xtrace feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > You can set this parameter to **true*	- only if the access log feature is enabled by specifying **AccessLogEnabled**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// The sampling rate of the Xtrace feature.
	//
	// Valid values: **1 to 10000**.
	//
	// > This parameter takes effect only if you set **TracingEnabled*	- to **true**.
	//
	// example:
	//
	// 100
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// The type of Xtrace. Set the value to **Zipkin**.
	//
	// > This parameter takes effect only if you set **TracingEnabled*	- to **true**.
	//
	// example:
	//
	// Zipkin
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s UpdateListenerLogConfigRequestAccessLogTracingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerLogConfigRequestAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) GetTracingEnabled() *bool {
	return s.TracingEnabled
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) GetTracingSample() *int32 {
	return s.TracingSample
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) GetTracingType() *string {
	return s.TracingType
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingEnabled(v bool) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingSample(v int32) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingType(v string) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) Validate() error {
	return dara.Validate(s)
}
