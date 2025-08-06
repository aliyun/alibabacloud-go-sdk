// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallRecordListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryCallRecordListResponseBodyData) *QueryCallRecordListResponseBody
	GetData() []*QueryCallRecordListResponseBodyData
	SetErrorCode(v string) *QueryCallRecordListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryCallRecordListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryCallRecordListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCallRecordListResponseBody
	GetSuccess() *bool
}

type QueryCallRecordListResponseBody struct {
	Data      []*QueryCallRecordListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCallRecordListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallRecordListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallRecordListResponseBody) GetData() []*QueryCallRecordListResponseBodyData {
	return s.Data
}

func (s *QueryCallRecordListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryCallRecordListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryCallRecordListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallRecordListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCallRecordListResponseBody) SetData(v []*QueryCallRecordListResponseBodyData) *QueryCallRecordListResponseBody {
	s.Data = v
	return s
}

func (s *QueryCallRecordListResponseBody) SetErrorCode(v string) *QueryCallRecordListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryCallRecordListResponseBody) SetErrorMsg(v string) *QueryCallRecordListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryCallRecordListResponseBody) SetRequestId(v string) *QueryCallRecordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallRecordListResponseBody) SetSuccess(v bool) *QueryCallRecordListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCallRecordListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCallRecordListResponseBodyData struct {
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryCallRecordListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCallRecordListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCallRecordListResponseBodyData) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryCallRecordListResponseBodyData) GetSignatureUrl() *string {
	return s.SignatureUrl
}

func (s *QueryCallRecordListResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCallRecordListResponseBodyData) SetDuration(v int32) *QueryCallRecordListResponseBodyData {
	s.Duration = &v
	return s
}

func (s *QueryCallRecordListResponseBodyData) SetSignatureUrl(v string) *QueryCallRecordListResponseBodyData {
	s.SignatureUrl = &v
	return s
}

func (s *QueryCallRecordListResponseBodyData) SetStartTime(v int64) *QueryCallRecordListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryCallRecordListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
