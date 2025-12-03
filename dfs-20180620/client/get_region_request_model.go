// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputRegionId(v string) *GetRegionRequest
	GetInputRegionId() *string
}

type GetRegionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s GetRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegionRequest) GoString() string {
	return s.String()
}

func (s *GetRegionRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *GetRegionRequest) SetInputRegionId(v string) *GetRegionRequest {
	s.InputRegionId = &v
	return s
}

func (s *GetRegionRequest) Validate() error {
	return dara.Validate(s)
}
