// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployAirflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RedeployAirflowResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *RedeployAirflowResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *RedeployAirflowResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *RedeployAirflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *RedeployAirflowResponseBody
	GetRequestId() *string
	SetRoot(v bool) *RedeployAirflowResponseBody
	GetRoot() *bool
	SetSuccess(v bool) *RedeployAirflowResponseBody
	GetSuccess() *bool
}

type RedeployAirflowResponseBody struct {
	// example:
	//
	// NOT_FOUND
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"Responses\\": {\\"responses\\": [{\\"Uuid\\": \\"af-63ly5o1yuf076ifrpbxpka396\\", \\"Success\\": True}]}}
	Root *bool `json:"Root,omitempty" xml:"Root,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RedeployAirflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedeployAirflowResponseBody) GoString() string {
	return s.String()
}

func (s *RedeployAirflowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RedeployAirflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RedeployAirflowResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RedeployAirflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RedeployAirflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedeployAirflowResponseBody) GetRoot() *bool {
	return s.Root
}

func (s *RedeployAirflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RedeployAirflowResponseBody) SetAccessDeniedDetail(v string) *RedeployAirflowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RedeployAirflowResponseBody) SetErrorCode(v string) *RedeployAirflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RedeployAirflowResponseBody) SetHttpStatusCode(v int64) *RedeployAirflowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RedeployAirflowResponseBody) SetMessage(v string) *RedeployAirflowResponseBody {
	s.Message = &v
	return s
}

func (s *RedeployAirflowResponseBody) SetRequestId(v string) *RedeployAirflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedeployAirflowResponseBody) SetRoot(v bool) *RedeployAirflowResponseBody {
	s.Root = &v
	return s
}

func (s *RedeployAirflowResponseBody) SetSuccess(v bool) *RedeployAirflowResponseBody {
	s.Success = &v
	return s
}

func (s *RedeployAirflowResponseBody) Validate() error {
	return dara.Validate(s)
}
