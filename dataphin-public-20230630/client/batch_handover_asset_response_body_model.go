// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchHandoverAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchHandoverAssetResponseBody
	GetCode() *string
	SetData(v *BatchHandoverAssetResponseBodyData) *BatchHandoverAssetResponseBody
	GetData() *BatchHandoverAssetResponseBodyData
	SetHttpStatusCode(v int32) *BatchHandoverAssetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchHandoverAssetResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchHandoverAssetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchHandoverAssetResponseBody
	GetSuccess() *bool
}

type BatchHandoverAssetResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BatchHandoverAssetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchHandoverAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchHandoverAssetResponseBody) GoString() string {
	return s.String()
}

func (s *BatchHandoverAssetResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchHandoverAssetResponseBody) GetData() *BatchHandoverAssetResponseBodyData {
	return s.Data
}

func (s *BatchHandoverAssetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchHandoverAssetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchHandoverAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchHandoverAssetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchHandoverAssetResponseBody) SetCode(v string) *BatchHandoverAssetResponseBody {
	s.Code = &v
	return s
}

func (s *BatchHandoverAssetResponseBody) SetData(v *BatchHandoverAssetResponseBodyData) *BatchHandoverAssetResponseBody {
	s.Data = v
	return s
}

func (s *BatchHandoverAssetResponseBody) SetHttpStatusCode(v int32) *BatchHandoverAssetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchHandoverAssetResponseBody) SetMessage(v string) *BatchHandoverAssetResponseBody {
	s.Message = &v
	return s
}

func (s *BatchHandoverAssetResponseBody) SetRequestId(v string) *BatchHandoverAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchHandoverAssetResponseBody) SetSuccess(v bool) *BatchHandoverAssetResponseBody {
	s.Success = &v
	return s
}

func (s *BatchHandoverAssetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchHandoverAssetResponseBodyData struct {
	// example:
	//
	// NullPointException
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0
	FailCount   *int32    `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	FailedGuids []*string `json:"FailedGuids,omitempty" xml:"FailedGuids,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BatchHandoverAssetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchHandoverAssetResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchHandoverAssetResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchHandoverAssetResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *BatchHandoverAssetResponseBodyData) GetFailedGuids() []*string {
	return s.FailedGuids
}

func (s *BatchHandoverAssetResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *BatchHandoverAssetResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchHandoverAssetResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *BatchHandoverAssetResponseBodyData) SetErrorMessage(v string) *BatchHandoverAssetResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *BatchHandoverAssetResponseBodyData) SetFailCount(v int32) *BatchHandoverAssetResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *BatchHandoverAssetResponseBodyData) SetFailedGuids(v []*string) *BatchHandoverAssetResponseBodyData {
	s.FailedGuids = v
	return s
}

func (s *BatchHandoverAssetResponseBodyData) SetStatus(v string) *BatchHandoverAssetResponseBodyData {
	s.Status = &v
	return s
}

func (s *BatchHandoverAssetResponseBodyData) SetSuccessCount(v int32) *BatchHandoverAssetResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *BatchHandoverAssetResponseBodyData) SetTotalCount(v int32) *BatchHandoverAssetResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *BatchHandoverAssetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
