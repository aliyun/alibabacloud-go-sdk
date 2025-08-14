// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarType(v string) *ListAvatarsRequest
	GetAvatarType() *string
	SetPageNo(v int32) *ListAvatarsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAvatarsRequest
	GetPageSize() *int32
}

type ListAvatarsRequest struct {
	// 	- The type of the digital human.
	//
	// 	- 2DAvatar
	//
	// example:
	//
	// 2DAvatar
	AvatarType *string `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	// 	- The page number.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 	- The number of entries per page.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAvatarsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarsRequest) GoString() string {
	return s.String()
}

func (s *ListAvatarsRequest) GetAvatarType() *string {
	return s.AvatarType
}

func (s *ListAvatarsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAvatarsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAvatarsRequest) SetAvatarType(v string) *ListAvatarsRequest {
	s.AvatarType = &v
	return s
}

func (s *ListAvatarsRequest) SetPageNo(v int32) *ListAvatarsRequest {
	s.PageNo = &v
	return s
}

func (s *ListAvatarsRequest) SetPageSize(v int32) *ListAvatarsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvatarsRequest) Validate() error {
	return dara.Validate(s)
}
