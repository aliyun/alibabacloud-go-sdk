// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApRunHistoryTimeSerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApRunHistoryTimeSerResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApRunHistoryTimeSerResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApRunHistoryTimeSerResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApRunHistoryTimeSerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetApRunHistoryTimeSerResponseBody
	GetRequestId() *string
}

type GetApRunHistoryTimeSerResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApRunHistoryTimeSerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApRunHistoryTimeSerResponseBody) GoString() string {
	return s.String()
}

func (s *GetApRunHistoryTimeSerResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApRunHistoryTimeSerResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApRunHistoryTimeSerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApRunHistoryTimeSerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApRunHistoryTimeSerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApRunHistoryTimeSerResponseBody) SetData(v map[string]interface{}) *GetApRunHistoryTimeSerResponseBody {
	s.Data = v
	return s
}

func (s *GetApRunHistoryTimeSerResponseBody) SetErrorCode(v int32) *GetApRunHistoryTimeSerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApRunHistoryTimeSerResponseBody) SetErrorMessage(v string) *GetApRunHistoryTimeSerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApRunHistoryTimeSerResponseBody) SetIsSuccess(v bool) *GetApRunHistoryTimeSerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApRunHistoryTimeSerResponseBody) SetRequestId(v string) *GetApRunHistoryTimeSerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApRunHistoryTimeSerResponseBody) Validate() error {
	return dara.Validate(s)
}
