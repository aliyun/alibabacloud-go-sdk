// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListCallbackRequest
	GetRegionId() *string
}

type ListCallbackRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallbackRequest) GoString() string {
	return s.String()
}

func (s *ListCallbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCallbackRequest) SetRegionId(v string) *ListCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *ListCallbackRequest) Validate() error {
	return dara.Validate(s)
}
