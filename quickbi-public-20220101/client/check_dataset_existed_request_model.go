// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDatasetExistedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *CheckDatasetExistedRequest
	GetCubeId() *string
}

type CheckDatasetExistedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s CheckDatasetExistedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDatasetExistedRequest) GoString() string {
	return s.String()
}

func (s *CheckDatasetExistedRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *CheckDatasetExistedRequest) SetCubeId(v string) *CheckDatasetExistedRequest {
	s.CubeId = &v
	return s
}

func (s *CheckDatasetExistedRequest) Validate() error {
	return dara.Validate(s)
}
