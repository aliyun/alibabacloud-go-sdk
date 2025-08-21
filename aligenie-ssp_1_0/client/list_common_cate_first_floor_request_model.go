// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateFirstFloorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *ListCommonCateFirstFloorRequest
	GetType() *string
}

type ListCommonCateFirstFloorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// song
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCommonCateFirstFloorRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateFirstFloorRequest) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorRequest) GetType() *string {
	return s.Type
}

func (s *ListCommonCateFirstFloorRequest) SetType(v string) *ListCommonCateFirstFloorRequest {
	s.Type = &v
	return s
}

func (s *ListCommonCateFirstFloorRequest) Validate() error {
	return dara.Validate(s)
}
