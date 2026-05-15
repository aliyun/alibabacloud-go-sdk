// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAirflowVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAirflowVersionsResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *ListAirflowVersionsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *ListAirflowVersionsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListAirflowVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAirflowVersionsResponseBody
	GetRequestId() *string
	SetRoot(v []*string) *ListAirflowVersionsResponseBody
	GetRoot() []*string
	SetSuccess(v bool) *ListAirflowVersionsResponseBody
	GetSuccess() *bool
}

type ListAirflowVersionsResponseBody struct {
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
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      []*string `json:"Root,omitempty" xml:"Root,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAirflowVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAirflowVersionsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAirflowVersionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAirflowVersionsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListAirflowVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAirflowVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAirflowVersionsResponseBody) GetRoot() []*string {
	return s.Root
}

func (s *ListAirflowVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAirflowVersionsResponseBody) SetAccessDeniedDetail(v string) *ListAirflowVersionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAirflowVersionsResponseBody) SetErrorCode(v string) *ListAirflowVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAirflowVersionsResponseBody) SetHttpStatusCode(v int64) *ListAirflowVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAirflowVersionsResponseBody) SetMessage(v string) *ListAirflowVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAirflowVersionsResponseBody) SetRequestId(v string) *ListAirflowVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAirflowVersionsResponseBody) SetRoot(v []*string) *ListAirflowVersionsResponseBody {
	s.Root = v
	return s
}

func (s *ListAirflowVersionsResponseBody) SetSuccess(v bool) *ListAirflowVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAirflowVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}
