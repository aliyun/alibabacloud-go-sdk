// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetSmartqStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *QueryDatasetSmartqStatusRequest
	GetCubeId() *string
}

type QueryDatasetSmartqStatusRequest struct {
	// Dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s QueryDatasetSmartqStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSmartqStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetSmartqStatusRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryDatasetSmartqStatusRequest) SetCubeId(v string) *QueryDatasetSmartqStatusRequest {
	s.CubeId = &v
	return s
}

func (s *QueryDatasetSmartqStatusRequest) Validate() error {
	return dara.Validate(s)
}
