// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *InvokeDiagnosisRequest
	GetChannel() *string
	SetParams(v string) *InvokeDiagnosisRequest
	GetParams() *string
	SetServiceName(v string) *InvokeDiagnosisRequest
	GetServiceName() *string
}

type InvokeDiagnosisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_assist
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "instance": "i-wz9gdv7qmdhusamc4dl01",
	//
	//     "uid": "xxxxxxxxxxxxxx",
	//
	//     "region": "cn-shenzhen"
	//
	// }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s InvokeDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisRequest) GetChannel() *string {
	return s.Channel
}

func (s *InvokeDiagnosisRequest) GetParams() *string {
	return s.Params
}

func (s *InvokeDiagnosisRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *InvokeDiagnosisRequest) SetChannel(v string) *InvokeDiagnosisRequest {
	s.Channel = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetParams(v string) *InvokeDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetServiceName(v string) *InvokeDiagnosisRequest {
	s.ServiceName = &v
	return s
}

func (s *InvokeDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
