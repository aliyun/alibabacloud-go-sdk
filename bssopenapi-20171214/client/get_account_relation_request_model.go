// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelationId(v int64) *GetAccountRelationRequest
	GetRelationId() *int64
	SetRequestId(v string) *GetAccountRelationRequest
	GetRequestId() *string
}

type GetAccountRelationRequest struct {
	// The account linking ID.
	//
	// example:
	//
	// 1234
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The unique request ID, used to identify the request and facilitate troubleshooting.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRelationRequest) GoString() string {
	return s.String()
}

func (s *GetAccountRelationRequest) GetRelationId() *int64 {
	return s.RelationId
}

func (s *GetAccountRelationRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountRelationRequest) SetRelationId(v int64) *GetAccountRelationRequest {
	s.RelationId = &v
	return s
}

func (s *GetAccountRelationRequest) SetRequestId(v string) *GetAccountRelationRequest {
	s.RequestId = &v
	return s
}

func (s *GetAccountRelationRequest) Validate() error {
	return dara.Validate(s)
}
