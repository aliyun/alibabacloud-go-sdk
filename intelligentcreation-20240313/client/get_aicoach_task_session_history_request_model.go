// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetAICoachTaskSessionHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetAICoachTaskSessionHistoryRequest
	GetPageSize() *int32
	SetSessionId(v string) *GetAICoachTaskSessionHistoryRequest
	GetSessionId() *string
	SetUid(v string) *GetAICoachTaskSessionHistoryRequest
	GetUid() *string
}

type GetAICoachTaskSessionHistoryRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 1251317954812712
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetAICoachTaskSessionHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAICoachTaskSessionHistoryRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAICoachTaskSessionHistoryRequest) GetUid() *string {
	return s.Uid
}

func (s *GetAICoachTaskSessionHistoryRequest) SetPageNumber(v int32) *GetAICoachTaskSessionHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryRequest) SetPageSize(v int32) *GetAICoachTaskSessionHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryRequest) SetSessionId(v string) *GetAICoachTaskSessionHistoryRequest {
	s.SessionId = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryRequest) SetUid(v string) *GetAICoachTaskSessionHistoryRequest {
	s.Uid = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryRequest) Validate() error {
	return dara.Validate(s)
}
