// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteGroupRequest
	GetId() *string
	SetOwnerId(v int64) *DeleteGroupRequest
	GetOwnerId() *int64
}

type DeleteGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33763950751395843
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetId() *string {
	return s.Id
}

func (s *DeleteGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteGroupRequest) SetId(v string) *DeleteGroupRequest {
	s.Id = &v
	return s
}

func (s *DeleteGroupRequest) SetOwnerId(v int64) *DeleteGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
