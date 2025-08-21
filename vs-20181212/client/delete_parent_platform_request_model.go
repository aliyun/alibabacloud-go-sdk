// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteParentPlatformRequest
	GetId() *string
	SetOwnerId(v int64) *DeleteParentPlatformRequest
	GetOwnerId() *int64
}

type DeleteParentPlatformRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 359*****374-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteParentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *DeleteParentPlatformRequest) GetId() *string {
	return s.Id
}

func (s *DeleteParentPlatformRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteParentPlatformRequest) SetId(v string) *DeleteParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *DeleteParentPlatformRequest) SetOwnerId(v int64) *DeleteParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteParentPlatformRequest) Validate() error {
	return dara.Validate(s)
}
