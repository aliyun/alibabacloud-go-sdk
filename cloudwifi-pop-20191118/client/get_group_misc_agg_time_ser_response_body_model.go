// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupMiscAggTimeSerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetGroupMiscAggTimeSerResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetGroupMiscAggTimeSerResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetGroupMiscAggTimeSerResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetGroupMiscAggTimeSerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetGroupMiscAggTimeSerResponseBody
	GetRequestId() *string
}

type GetGroupMiscAggTimeSerResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupMiscAggTimeSerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupMiscAggTimeSerResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupMiscAggTimeSerResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetGroupMiscAggTimeSerResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetGroupMiscAggTimeSerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetGroupMiscAggTimeSerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetGroupMiscAggTimeSerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupMiscAggTimeSerResponseBody) SetData(v map[string]interface{}) *GetGroupMiscAggTimeSerResponseBody {
	s.Data = v
	return s
}

func (s *GetGroupMiscAggTimeSerResponseBody) SetErrorCode(v int32) *GetGroupMiscAggTimeSerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGroupMiscAggTimeSerResponseBody) SetErrorMessage(v string) *GetGroupMiscAggTimeSerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGroupMiscAggTimeSerResponseBody) SetIsSuccess(v bool) *GetGroupMiscAggTimeSerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetGroupMiscAggTimeSerResponseBody) SetRequestId(v string) *GetGroupMiscAggTimeSerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupMiscAggTimeSerResponseBody) Validate() error {
	return dara.Validate(s)
}
