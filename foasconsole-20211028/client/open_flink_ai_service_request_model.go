// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenFlinkAiServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *OpenFlinkAiServiceRequest
	GetRegion() *string
}

type OpenFlinkAiServiceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s OpenFlinkAiServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenFlinkAiServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenFlinkAiServiceRequest) GetRegion() *string {
	return s.Region
}

func (s *OpenFlinkAiServiceRequest) SetRegion(v string) *OpenFlinkAiServiceRequest {
	s.Region = &v
	return s
}

func (s *OpenFlinkAiServiceRequest) Validate() error {
	return dara.Validate(s)
}
