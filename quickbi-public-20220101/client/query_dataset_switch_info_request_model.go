// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetSwitchInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *QueryDatasetSwitchInfoRequest
	GetCubeId() *string
}

type QueryDatasetSwitchInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s QueryDatasetSwitchInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSwitchInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryDatasetSwitchInfoRequest) SetCubeId(v string) *QueryDatasetSwitchInfoRequest {
	s.CubeId = &v
	return s
}

func (s *QueryDatasetSwitchInfoRequest) Validate() error {
	return dara.Validate(s)
}
