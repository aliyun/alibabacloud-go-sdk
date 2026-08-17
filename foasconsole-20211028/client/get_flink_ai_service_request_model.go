// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlinkAiServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GetFlinkAiServiceRequest
	GetRegion() *string
}

type GetFlinkAiServiceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetFlinkAiServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceRequest) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceRequest) GetRegion() *string {
	return s.Region
}

func (s *GetFlinkAiServiceRequest) SetRegion(v string) *GetFlinkAiServiceRequest {
	s.Region = &v
	return s
}

func (s *GetFlinkAiServiceRequest) Validate() error {
	return dara.Validate(s)
}
