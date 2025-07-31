// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBatchTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBatchTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBatchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBatchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBatchTaskResponseBody
	GetSuccess() *bool
	SetUpdateResult(v *UpdateBatchTaskResponseBodyUpdateResult) *UpdateBatchTaskResponseBody
	GetUpdateResult() *UpdateBatchTaskResponseBodyUpdateResult
}

type UpdateBatchTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	UpdateResult *UpdateBatchTaskResponseBodyUpdateResult `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty" type:"Struct"`
}

func (s UpdateBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBatchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBatchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBatchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBatchTaskResponseBody) GetUpdateResult() *UpdateBatchTaskResponseBodyUpdateResult {
	return s.UpdateResult
}

func (s *UpdateBatchTaskResponseBody) SetCode(v string) *UpdateBatchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBatchTaskResponseBody) SetHttpStatusCode(v int32) *UpdateBatchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBatchTaskResponseBody) SetMessage(v string) *UpdateBatchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBatchTaskResponseBody) SetRequestId(v string) *UpdateBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBatchTaskResponseBody) SetSuccess(v bool) *UpdateBatchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBatchTaskResponseBody) SetUpdateResult(v *UpdateBatchTaskResponseBodyUpdateResult) *UpdateBatchTaskResponseBody {
	s.UpdateResult = v
	return s
}

func (s *UpdateBatchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateBatchTaskResponseBodyUpdateResult struct {
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s UpdateBatchTaskResponseBodyUpdateResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskResponseBodyUpdateResult) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskResponseBodyUpdateResult) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateBatchTaskResponseBodyUpdateResult) SetFileId(v int64) *UpdateBatchTaskResponseBodyUpdateResult {
	s.FileId = &v
	return s
}

func (s *UpdateBatchTaskResponseBodyUpdateResult) Validate() error {
	return dara.Validate(s)
}
