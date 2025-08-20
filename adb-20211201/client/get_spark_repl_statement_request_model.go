// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkReplStatementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkReplStatementRequest
	GetAppId() *string
	SetSessionId(v int64) *GetSparkReplStatementRequest
	GetSessionId() *int64
	SetStatementId(v int64) *GetSparkReplStatementRequest
	GetStatementId() *int64
}

type GetSparkReplStatementRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query Spark application IDs.
	//
	// example:
	//
	// s202411071444hzdvk486d9d200****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the session that executes the code.
	//
	// example:
	//
	// 1
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The unique ID of the code block in the Spark job.
	//
	// example:
	//
	// 123
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s GetSparkReplStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplStatementRequest) GoString() string {
	return s.String()
}

func (s *GetSparkReplStatementRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkReplStatementRequest) GetSessionId() *int64 {
	return s.SessionId
}

func (s *GetSparkReplStatementRequest) GetStatementId() *int64 {
	return s.StatementId
}

func (s *GetSparkReplStatementRequest) SetAppId(v string) *GetSparkReplStatementRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkReplStatementRequest) SetSessionId(v int64) *GetSparkReplStatementRequest {
	s.SessionId = &v
	return s
}

func (s *GetSparkReplStatementRequest) SetStatementId(v int64) *GetSparkReplStatementRequest {
	s.StatementId = &v
	return s
}

func (s *GetSparkReplStatementRequest) Validate() error {
	return dara.Validate(s)
}
