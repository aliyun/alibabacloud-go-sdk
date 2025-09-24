// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteFaceGroupRequest
	GetId() *string
}

type DeleteFaceGroupRequest struct {
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 16681
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteFaceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupRequest) GetId() *string {
	return s.Id
}

func (s *DeleteFaceGroupRequest) SetId(v string) *DeleteFaceGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteFaceGroupRequest) Validate() error {
	return dara.Validate(s)
}
