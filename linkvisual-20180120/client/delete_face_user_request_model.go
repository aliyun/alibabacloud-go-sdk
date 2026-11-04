// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *DeleteFaceUserRequest
	GetIsolationId() *string
	SetUserId(v string) *DeleteFaceUserRequest
	GetUserId() *string
}

type DeleteFaceUserRequest struct {
	// The business isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHzXiHu
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The unique internal user ID of the face service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3mtjaqebeq4dulcy
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteFaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *DeleteFaceUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteFaceUserRequest) SetIsolationId(v string) *DeleteFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserRequest) SetUserId(v string) *DeleteFaceUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteFaceUserRequest) Validate() error {
	return dara.Validate(s)
}
