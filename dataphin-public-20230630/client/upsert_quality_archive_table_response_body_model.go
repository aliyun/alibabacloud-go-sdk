// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityArchiveTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpsertQualityArchiveTableResponseBody
	GetCode() *string
	SetData(v *UpsertQualityArchiveTableResponseBodyData) *UpsertQualityArchiveTableResponseBody
	GetData() *UpsertQualityArchiveTableResponseBodyData
	SetHttpStatusCode(v int32) *UpsertQualityArchiveTableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpsertQualityArchiveTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertQualityArchiveTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertQualityArchiveTableResponseBody
	GetSuccess() *bool
}

type UpsertQualityArchiveTableResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submit status of the asynchronous task.
	Data *UpsertQualityArchiveTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error details returned by the backend.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpsertQualityArchiveTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityArchiveTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertQualityArchiveTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpsertQualityArchiveTableResponseBody) GetData() *UpsertQualityArchiveTableResponseBodyData {
	return s.Data
}

func (s *UpsertQualityArchiveTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpsertQualityArchiveTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertQualityArchiveTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertQualityArchiveTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertQualityArchiveTableResponseBody) SetCode(v string) *UpsertQualityArchiveTableResponseBody {
	s.Code = &v
	return s
}

func (s *UpsertQualityArchiveTableResponseBody) SetData(v *UpsertQualityArchiveTableResponseBodyData) *UpsertQualityArchiveTableResponseBody {
	s.Data = v
	return s
}

func (s *UpsertQualityArchiveTableResponseBody) SetHttpStatusCode(v int32) *UpsertQualityArchiveTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpsertQualityArchiveTableResponseBody) SetMessage(v string) *UpsertQualityArchiveTableResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertQualityArchiveTableResponseBody) SetRequestId(v string) *UpsertQualityArchiveTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertQualityArchiveTableResponseBody) SetSuccess(v bool) *UpsertQualityArchiveTableResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertQualityArchiveTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityArchiveTableResponseBodyData struct {
	// The progress ID of the asynchronous task. This is an asynchronous operation. After a successful submission, only this ID is returned. Call GetQualityArchiveTableProgress to poll the task status and retrieve the final archived table information.
	//
	// example:
	//
	// d78f0b5c9a1e4f2ab3c6d5e4f7a8b9c0
	ProgressId *string `json:"ProgressId,omitempty" xml:"ProgressId,omitempty"`
}

func (s UpsertQualityArchiveTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityArchiveTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpsertQualityArchiveTableResponseBodyData) GetProgressId() *string {
	return s.ProgressId
}

func (s *UpsertQualityArchiveTableResponseBodyData) SetProgressId(v string) *UpsertQualityArchiveTableResponseBodyData {
	s.ProgressId = &v
	return s
}

func (s *UpsertQualityArchiveTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
