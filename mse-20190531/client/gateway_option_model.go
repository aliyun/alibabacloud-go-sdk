// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayOption interface {
	dara.Model
	String() string
	GoString() string
	SetDisableHttp2Alpn(v bool) *GatewayOption
	GetDisableHttp2Alpn() *bool
	SetEnableHardwareAcceleration(v bool) *GatewayOption
	GetEnableHardwareAcceleration() *bool
	SetEnableWaf(v bool) *GatewayOption
	GetEnableWaf() *bool
	SetLogConfigDetails(v *GatewayOptionLogConfigDetails) *GatewayOption
	GetLogConfigDetails() *GatewayOptionLogConfigDetails
	SetTraceDetails(v *GatewayOptionTraceDetails) *GatewayOption
	GetTraceDetails() *GatewayOptionTraceDetails
}

type GatewayOption struct {
	// Specifies whether to disable the HTTP/2 protocol.
	//
	// example:
	//
	// true
	DisableHttp2Alpn *bool `json:"DisableHttp2Alpn,omitempty" xml:"DisableHttp2Alpn,omitempty"`
	// Specifies whether to enable hardware acceleration.
	//
	// example:
	//
	// true
	EnableHardwareAcceleration *bool `json:"EnableHardwareAcceleration,omitempty" xml:"EnableHardwareAcceleration,omitempty"`
	// Specifies whether to enable Web Application Firewall (WAF).
	//
	// example:
	//
	// true
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The description of Simple Log Service.
	LogConfigDetails *GatewayOptionLogConfigDetails `json:"LogConfigDetails,omitempty" xml:"LogConfigDetails,omitempty" type:"Struct"`
	// The data structure.
	TraceDetails *GatewayOptionTraceDetails `json:"TraceDetails,omitempty" xml:"TraceDetails,omitempty" type:"Struct"`
}

func (s GatewayOption) String() string {
	return dara.Prettify(s)
}

func (s GatewayOption) GoString() string {
	return s.String()
}

func (s *GatewayOption) GetDisableHttp2Alpn() *bool {
	return s.DisableHttp2Alpn
}

func (s *GatewayOption) GetEnableHardwareAcceleration() *bool {
	return s.EnableHardwareAcceleration
}

func (s *GatewayOption) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *GatewayOption) GetLogConfigDetails() *GatewayOptionLogConfigDetails {
	return s.LogConfigDetails
}

func (s *GatewayOption) GetTraceDetails() *GatewayOptionTraceDetails {
	return s.TraceDetails
}

func (s *GatewayOption) SetDisableHttp2Alpn(v bool) *GatewayOption {
	s.DisableHttp2Alpn = &v
	return s
}

func (s *GatewayOption) SetEnableHardwareAcceleration(v bool) *GatewayOption {
	s.EnableHardwareAcceleration = &v
	return s
}

func (s *GatewayOption) SetEnableWaf(v bool) *GatewayOption {
	s.EnableWaf = &v
	return s
}

func (s *GatewayOption) SetLogConfigDetails(v *GatewayOptionLogConfigDetails) *GatewayOption {
	s.LogConfigDetails = v
	return s
}

func (s *GatewayOption) SetTraceDetails(v *GatewayOptionTraceDetails) *GatewayOption {
	s.TraceDetails = v
	return s
}

func (s *GatewayOption) Validate() error {
	if s.LogConfigDetails != nil {
		if err := s.LogConfigDetails.Validate(); err != nil {
			return err
		}
	}
	if s.TraceDetails != nil {
		if err := s.TraceDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GatewayOptionLogConfigDetails struct {
	// Specifies whether to activate Simple Log Service.
	//
	// Valid value:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	LogEnabled *bool `json:"LogEnabled,omitempty" xml:"LogEnabled,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// name
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the destination Simple Log Service project.
	//
	// example:
	//
	// project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GatewayOptionLogConfigDetails) String() string {
	return dara.Prettify(s)
}

func (s GatewayOptionLogConfigDetails) GoString() string {
	return s.String()
}

func (s *GatewayOptionLogConfigDetails) GetLogEnabled() *bool {
	return s.LogEnabled
}

func (s *GatewayOptionLogConfigDetails) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GatewayOptionLogConfigDetails) GetProjectName() *string {
	return s.ProjectName
}

func (s *GatewayOptionLogConfigDetails) SetLogEnabled(v bool) *GatewayOptionLogConfigDetails {
	s.LogEnabled = &v
	return s
}

func (s *GatewayOptionLogConfigDetails) SetLogStoreName(v string) *GatewayOptionLogConfigDetails {
	s.LogStoreName = &v
	return s
}

func (s *GatewayOptionLogConfigDetails) SetProjectName(v string) *GatewayOptionLogConfigDetails {
	s.ProjectName = &v
	return s
}

func (s *GatewayOptionLogConfigDetails) Validate() error {
	return dara.Validate(s)
}

type GatewayOptionTraceDetails struct {
	// The sampling rate of Tracing Analysis.
	//
	// example:
	//
	// 10
	Sample *int64 `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The ID of the SkyWalking service. This parameter is required if TraceType is set to SKYWALKING.
	//
	// example:
	//
	// 10458
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The port of the SkyWalking service. This parameter is required if TraceType is set to SKYWALKING.
	//
	// example:
	//
	// 80
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// Specifies whether to activate Tracing Analysis.
	//
	// Valid value:
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	TraceEnabled *bool `json:"TraceEnabled,omitempty" xml:"TraceEnabled,omitempty"`
	// The type of Tracing Analysis. Valid values: XTRACE and SKYWALKING.
	//
	// example:
	//
	// XTRACE
	TraceType *string `json:"TraceType,omitempty" xml:"TraceType,omitempty"`
}

func (s GatewayOptionTraceDetails) String() string {
	return dara.Prettify(s)
}

func (s GatewayOptionTraceDetails) GoString() string {
	return s.String()
}

func (s *GatewayOptionTraceDetails) GetSample() *int64 {
	return s.Sample
}

func (s *GatewayOptionTraceDetails) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GatewayOptionTraceDetails) GetServicePort() *string {
	return s.ServicePort
}

func (s *GatewayOptionTraceDetails) GetTraceEnabled() *bool {
	return s.TraceEnabled
}

func (s *GatewayOptionTraceDetails) GetTraceType() *string {
	return s.TraceType
}

func (s *GatewayOptionTraceDetails) SetSample(v int64) *GatewayOptionTraceDetails {
	s.Sample = &v
	return s
}

func (s *GatewayOptionTraceDetails) SetServiceId(v int64) *GatewayOptionTraceDetails {
	s.ServiceId = &v
	return s
}

func (s *GatewayOptionTraceDetails) SetServicePort(v string) *GatewayOptionTraceDetails {
	s.ServicePort = &v
	return s
}

func (s *GatewayOptionTraceDetails) SetTraceEnabled(v bool) *GatewayOptionTraceDetails {
	s.TraceEnabled = &v
	return s
}

func (s *GatewayOptionTraceDetails) SetTraceType(v string) *GatewayOptionTraceDetails {
	s.TraceType = &v
	return s
}

func (s *GatewayOptionTraceDetails) Validate() error {
	return dara.Validate(s)
}
