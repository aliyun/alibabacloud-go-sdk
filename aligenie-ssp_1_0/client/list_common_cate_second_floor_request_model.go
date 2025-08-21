// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateSecondFloorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentCateId(v int64) *ListCommonCateSecondFloorRequest
	GetParentCateId() *int64
}

type ListCommonCateSecondFloorRequest struct {
	// example:
	//
	// 80010
	ParentCateId *int64 `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCommonCateSecondFloorRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateSecondFloorRequest) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorRequest) GetParentCateId() *int64 {
	return s.ParentCateId
}

func (s *ListCommonCateSecondFloorRequest) SetParentCateId(v int64) *ListCommonCateSecondFloorRequest {
	s.ParentCateId = &v
	return s
}

func (s *ListCommonCateSecondFloorRequest) Validate() error {
	return dara.Validate(s)
}
