// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteDirectoryRequest
	GetId() *string
	SetOwnerId(v int64) *DeleteDirectoryRequest
	GetOwnerId() *int64
}

type DeleteDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryRequest) GetId() *string {
	return s.Id
}

func (s *DeleteDirectoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDirectoryRequest) SetId(v string) *DeleteDirectoryRequest {
	s.Id = &v
	return s
}

func (s *DeleteDirectoryRequest) SetOwnerId(v int64) *DeleteDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
