// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContacterId(v int64) *DeleteContacterRequest
	GetContacterId() *int64
	SetUserId(v int64) *DeleteContacterRequest
	GetUserId() *int64
}

type DeleteContacterRequest struct {
	// This parameter is required.
	ContacterId *int64 `json:"ContacterId,omitempty" xml:"ContacterId,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteContacterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContacterRequest) GoString() string {
	return s.String()
}

func (s *DeleteContacterRequest) GetContacterId() *int64 {
	return s.ContacterId
}

func (s *DeleteContacterRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DeleteContacterRequest) SetContacterId(v int64) *DeleteContacterRequest {
	s.ContacterId = &v
	return s
}

func (s *DeleteContacterRequest) SetUserId(v int64) *DeleteContacterRequest {
	s.UserId = &v
	return s
}

func (s *DeleteContacterRequest) Validate() error {
	return dara.Validate(s)
}
