// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *CancelSqlPreviewRequest
	GetQueryId() *string
	SetSessionClusterName(v string) *CancelSqlPreviewRequest
	GetSessionClusterName() *string
	SetSessionId(v string) *CancelSqlPreviewRequest
	GetSessionId() *string
}

type CancelSqlPreviewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 19426537348647121698828223472
	QueryId *string `json:"queryId,omitempty" xml:"queryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-session
	SessionClusterName *string `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a737f33bbdb7419db9ee8037bb51e73a
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CancelSqlPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPreviewRequest) GoString() string {
	return s.String()
}

func (s *CancelSqlPreviewRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *CancelSqlPreviewRequest) GetSessionClusterName() *string {
	return s.SessionClusterName
}

func (s *CancelSqlPreviewRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CancelSqlPreviewRequest) SetQueryId(v string) *CancelSqlPreviewRequest {
	s.QueryId = &v
	return s
}

func (s *CancelSqlPreviewRequest) SetSessionClusterName(v string) *CancelSqlPreviewRequest {
	s.SessionClusterName = &v
	return s
}

func (s *CancelSqlPreviewRequest) SetSessionId(v string) *CancelSqlPreviewRequest {
	s.SessionId = &v
	return s
}

func (s *CancelSqlPreviewRequest) Validate() error {
	return dara.Validate(s)
}
