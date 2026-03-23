// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRadioRunHistoryTimeSerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetRadioRunHistoryTimeSerResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetRadioRunHistoryTimeSerResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetRadioRunHistoryTimeSerResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetRadioRunHistoryTimeSerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetRadioRunHistoryTimeSerResponseBody
	GetRequestId() *string
}

type GetRadioRunHistoryTimeSerResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRadioRunHistoryTimeSerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRadioRunHistoryTimeSerResponseBody) GoString() string {
	return s.String()
}

func (s *GetRadioRunHistoryTimeSerResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetRadioRunHistoryTimeSerResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetRadioRunHistoryTimeSerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRadioRunHistoryTimeSerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRadioRunHistoryTimeSerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRadioRunHistoryTimeSerResponseBody) SetData(v map[string]interface{}) *GetRadioRunHistoryTimeSerResponseBody {
	s.Data = v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponseBody) SetErrorCode(v int32) *GetRadioRunHistoryTimeSerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponseBody) SetErrorMessage(v string) *GetRadioRunHistoryTimeSerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponseBody) SetIsSuccess(v bool) *GetRadioRunHistoryTimeSerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponseBody) SetRequestId(v string) *GetRadioRunHistoryTimeSerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerResponseBody) Validate() error {
	return dara.Validate(s)
}
