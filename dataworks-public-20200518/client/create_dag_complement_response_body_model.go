// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDagComplementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *CreateDagComplementResponseBody
	GetData() []*int64
	SetErrorCode(v string) *CreateDagComplementResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDagComplementResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateDagComplementResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateDagComplementResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDagComplementResponseBody
	GetSuccess() *bool
}

type CreateDagComplementResponseBody struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s CreateDagComplementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDagComplementResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDagComplementResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *CreateDagComplementResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDagComplementResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDagComplementResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDagComplementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDagComplementResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDagComplementResponseBody) SetData(v []*int64) *CreateDagComplementResponseBody {
	s.Data = v
	return s
}

func (s *CreateDagComplementResponseBody) SetErrorCode(v string) *CreateDagComplementResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDagComplementResponseBody) SetErrorMessage(v string) *CreateDagComplementResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDagComplementResponseBody) SetHttpStatusCode(v int32) *CreateDagComplementResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDagComplementResponseBody) SetRequestId(v string) *CreateDagComplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDagComplementResponseBody) SetSuccess(v bool) *CreateDagComplementResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDagComplementResponseBody) Validate() error {
	return dara.Validate(s)
}
