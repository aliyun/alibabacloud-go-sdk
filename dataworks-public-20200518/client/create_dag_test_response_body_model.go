// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDagTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateDagTestResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateDagTestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDagTestResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateDagTestResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateDagTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDagTestResponseBody
	GetSuccess() *bool
}

type CreateDagTestResponseBody struct {
	// example:
	//
	// 3333333
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDagTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDagTestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDagTestResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDagTestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDagTestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDagTestResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDagTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDagTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDagTestResponseBody) SetData(v int64) *CreateDagTestResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDagTestResponseBody) SetErrorCode(v string) *CreateDagTestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDagTestResponseBody) SetErrorMessage(v string) *CreateDagTestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDagTestResponseBody) SetHttpStatusCode(v int32) *CreateDagTestResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDagTestResponseBody) SetRequestId(v string) *CreateDagTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDagTestResponseBody) SetSuccess(v bool) *CreateDagTestResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDagTestResponseBody) Validate() error {
	return dara.Validate(s)
}
