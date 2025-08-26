// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAirflowResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *DeleteAirflowResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *DeleteAirflowResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DeleteAirflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAirflowResponseBody
	GetRequestId() *string
	SetRoot(v *DeleteAirflowResponseBodyRoot) *DeleteAirflowResponseBody
	GetRoot() *DeleteAirflowResponseBodyRoot
	SetSuccess(v bool) *DeleteAirflowResponseBody
	GetSuccess() *bool
}

type DeleteAirflowResponseBody struct {
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
	// Instance not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-****
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *DeleteAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAirflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAirflowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAirflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAirflowResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DeleteAirflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAirflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAirflowResponseBody) GetRoot() *DeleteAirflowResponseBodyRoot {
	return s.Root
}

func (s *DeleteAirflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAirflowResponseBody) SetAccessDeniedDetail(v string) *DeleteAirflowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAirflowResponseBody) SetErrorCode(v string) *DeleteAirflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAirflowResponseBody) SetHttpStatusCode(v int64) *DeleteAirflowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAirflowResponseBody) SetMessage(v string) *DeleteAirflowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAirflowResponseBody) SetRequestId(v string) *DeleteAirflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAirflowResponseBody) SetRoot(v *DeleteAirflowResponseBodyRoot) *DeleteAirflowResponseBody {
	s.Root = v
	return s
}

func (s *DeleteAirflowResponseBody) SetSuccess(v bool) *DeleteAirflowResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAirflowResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteAirflowResponseBodyRoot struct {
	Responses []*DeleteAirflowResponseBodyRootResponses `json:"Responses,omitempty" xml:"Responses,omitempty" type:"Repeated"`
}

func (s DeleteAirflowResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirflowResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *DeleteAirflowResponseBodyRoot) GetResponses() []*DeleteAirflowResponseBodyRootResponses {
	return s.Responses
}

func (s *DeleteAirflowResponseBodyRoot) SetResponses(v []*DeleteAirflowResponseBodyRootResponses) *DeleteAirflowResponseBodyRoot {
	s.Responses = v
	return s
}

func (s *DeleteAirflowResponseBodyRoot) Validate() error {
	return dara.Validate(s)
}

type DeleteAirflowResponseBodyRootResponses struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// af-test****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DeleteAirflowResponseBodyRootResponses) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirflowResponseBodyRootResponses) GoString() string {
	return s.String()
}

func (s *DeleteAirflowResponseBodyRootResponses) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAirflowResponseBodyRootResponses) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteAirflowResponseBodyRootResponses) SetSuccess(v bool) *DeleteAirflowResponseBodyRootResponses {
	s.Success = &v
	return s
}

func (s *DeleteAirflowResponseBodyRootResponses) SetUuid(v string) *DeleteAirflowResponseBodyRootResponses {
	s.Uuid = &v
	return s
}

func (s *DeleteAirflowResponseBodyRootResponses) Validate() error {
	return dara.Validate(s)
}
