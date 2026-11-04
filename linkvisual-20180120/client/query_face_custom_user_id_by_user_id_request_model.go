// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceCustomUserIdByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *QueryFaceCustomUserIdByUserIdRequest
	GetIsolationId() *string
	SetUserId(v string) *QueryFaceCustomUserIdByUserIdRequest
	GetUserId() *string
}

type QueryFaceCustomUserIdByUserIdRequest struct {
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

func (s QueryFaceCustomUserIdByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceCustomUserIdByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceCustomUserIdByUserIdRequest) SetIsolationId(v string) *QueryFaceCustomUserIdByUserIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdRequest) SetUserId(v string) *QueryFaceCustomUserIdByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
