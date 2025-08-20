// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSparkReplStatementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CancelSparkReplStatementRequest
	GetAppId() *string
	SetSessionId(v int64) *CancelSparkReplStatementRequest
	GetSessionId() *int64
	SetStatementId(v int64) *CancelSparkReplStatementRequest
	GetStatementId() *int64
}

type CancelSparkReplStatementRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query Spark application IDs.
	//
	// example:
	//
	// s202411071444hzdvk486d9d2001****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 456
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The unique ID of the code block in the Spark job.
	//
	// example:
	//
	// 123
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s CancelSparkReplStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkReplStatementRequest) GoString() string {
	return s.String()
}

func (s *CancelSparkReplStatementRequest) GetAppId() *string {
	return s.AppId
}

func (s *CancelSparkReplStatementRequest) GetSessionId() *int64 {
	return s.SessionId
}

func (s *CancelSparkReplStatementRequest) GetStatementId() *int64 {
	return s.StatementId
}

func (s *CancelSparkReplStatementRequest) SetAppId(v string) *CancelSparkReplStatementRequest {
	s.AppId = &v
	return s
}

func (s *CancelSparkReplStatementRequest) SetSessionId(v int64) *CancelSparkReplStatementRequest {
	s.SessionId = &v
	return s
}

func (s *CancelSparkReplStatementRequest) SetStatementId(v int64) *CancelSparkReplStatementRequest {
	s.StatementId = &v
	return s
}

func (s *CancelSparkReplStatementRequest) Validate() error {
	return dara.Validate(s)
}
