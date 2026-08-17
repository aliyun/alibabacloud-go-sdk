// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseFlinkAiServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *CloseFlinkAiServiceRequest
	GetRegion() *string
}

type CloseFlinkAiServiceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CloseFlinkAiServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseFlinkAiServiceRequest) GoString() string {
	return s.String()
}

func (s *CloseFlinkAiServiceRequest) GetRegion() *string {
	return s.Region
}

func (s *CloseFlinkAiServiceRequest) SetRegion(v string) *CloseFlinkAiServiceRequest {
	s.Region = &v
	return s
}

func (s *CloseFlinkAiServiceRequest) Validate() error {
	return dara.Validate(s)
}
