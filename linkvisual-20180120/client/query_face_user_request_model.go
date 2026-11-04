// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceUserRequest
	GetIsolationId() *string
	SetUserId(v string) *QueryFaceUserRequest
	GetUserId() *string
}

type QueryFaceUserRequest struct {
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceUserRequest) SetIsolationId(v string) *QueryFaceUserRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserRequest) SetUserId(v string) *QueryFaceUserRequest {
	s.UserId = &v
	return s
}

func (s *QueryFaceUserRequest) Validate() error {
	return dara.Validate(s)
}
