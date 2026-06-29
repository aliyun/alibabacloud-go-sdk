// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBatchTaskResponseBody
	GetCode() *string
	SetCreateResult(v *CreateBatchTaskResponseBodyCreateResult) *CreateBatchTaskResponseBody
	GetCreateResult() *CreateBatchTaskResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateBatchTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBatchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBatchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBatchTaskResponseBody
	GetSuccess() *bool
}

type CreateBatchTaskResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation result.
	CreateResult *CreateBatchTaskResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBatchTaskResponseBody) GetCreateResult() *CreateBatchTaskResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateBatchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBatchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBatchTaskResponseBody) SetCode(v string) *CreateBatchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBatchTaskResponseBody) SetCreateResult(v *CreateBatchTaskResponseBodyCreateResult) *CreateBatchTaskResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateBatchTaskResponseBody) SetHttpStatusCode(v int32) *CreateBatchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBatchTaskResponseBody) SetMessage(v string) *CreateBatchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBatchTaskResponseBody) SetRequestId(v string) *CreateBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchTaskResponseBody) SetSuccess(v bool) *CreateBatchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBatchTaskResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchTaskResponseBodyCreateResult struct {
	// The node ID in the node directory tree.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s CreateBatchTaskResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskResponseBodyCreateResult) GetFileId() *int64 {
	return s.FileId
}

func (s *CreateBatchTaskResponseBodyCreateResult) SetFileId(v int64) *CreateBatchTaskResponseBodyCreateResult {
	s.FileId = &v
	return s
}

func (s *CreateBatchTaskResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
