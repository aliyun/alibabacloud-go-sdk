// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAifgAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *StartAifgAnalysisRequest
	GetChannel() *string
	SetParams(v string) *StartAifgAnalysisRequest
	GetParams() *string
	SetRegion(v string) *StartAifgAnalysisRequest
	GetRegion() *string
	SetServiceName(v string) *StartAifgAnalysisRequest
	GetServiceName() *string
}

type StartAifgAnalysisRequest struct {
	// The channel type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The additional information.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"train_taskid": "train-92abd24", "duration": "10"}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// The region information.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// slowNode
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s StartAifgAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAifgAnalysisRequest) GoString() string {
	return s.String()
}

func (s *StartAifgAnalysisRequest) GetChannel() *string {
	return s.Channel
}

func (s *StartAifgAnalysisRequest) GetParams() *string {
	return s.Params
}

func (s *StartAifgAnalysisRequest) GetRegion() *string {
	return s.Region
}

func (s *StartAifgAnalysisRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *StartAifgAnalysisRequest) SetChannel(v string) *StartAifgAnalysisRequest {
	s.Channel = &v
	return s
}

func (s *StartAifgAnalysisRequest) SetParams(v string) *StartAifgAnalysisRequest {
	s.Params = &v
	return s
}

func (s *StartAifgAnalysisRequest) SetRegion(v string) *StartAifgAnalysisRequest {
	s.Region = &v
	return s
}

func (s *StartAifgAnalysisRequest) SetServiceName(v string) *StartAifgAnalysisRequest {
	s.ServiceName = &v
	return s
}

func (s *StartAifgAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
