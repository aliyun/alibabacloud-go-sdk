// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkReplSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkReplSessionRequest
	GetAppId() *string
	SetSessionId(v int64) *GetSparkReplSessionRequest
	GetSessionId() *int64
}

type GetSparkReplSessionRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query all application IDs.
	//
	// example:
	//
	// s202411071444hzdvk486d9d200****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the session that executes the code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetSparkReplSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplSessionRequest) GoString() string {
	return s.String()
}

func (s *GetSparkReplSessionRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkReplSessionRequest) GetSessionId() *int64 {
	return s.SessionId
}

func (s *GetSparkReplSessionRequest) SetAppId(v string) *GetSparkReplSessionRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkReplSessionRequest) SetSessionId(v int64) *GetSparkReplSessionRequest {
	s.SessionId = &v
	return s
}

func (s *GetSparkReplSessionRequest) Validate() error {
	return dara.Validate(s)
}
