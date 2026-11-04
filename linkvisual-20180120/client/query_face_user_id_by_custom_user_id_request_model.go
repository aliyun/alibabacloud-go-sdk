// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserIdByCustomUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdRequest
	GetCustomUserId() *string
	SetIsolationId(v string) *QueryFaceUserIdByCustomUserIdRequest
	GetIsolationId() *string
}

type QueryFaceUserIdByCustomUserIdRequest struct {
	// The user ID in the caller system.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestCustomUserId
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdRequest) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *QueryFaceUserIdByCustomUserIdRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetCustomUserId(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdRequest) SetIsolationId(v string) *QueryFaceUserIdByCustomUserIdRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdRequest) Validate() error {
	return dara.Validate(s)
}
