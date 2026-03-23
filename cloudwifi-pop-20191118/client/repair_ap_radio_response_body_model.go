// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepairApRadioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *RepairApRadioResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v bool) *RepairApRadioResponseBody
	GetErrorCode() *bool
	SetErrorMessage(v string) *RepairApRadioResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *RepairApRadioResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *RepairApRadioResponseBody
	GetRequestId() *string
}

type RepairApRadioResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *bool                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RepairApRadioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RepairApRadioResponseBody) GoString() string {
	return s.String()
}

func (s *RepairApRadioResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *RepairApRadioResponseBody) GetErrorCode() *bool {
	return s.ErrorCode
}

func (s *RepairApRadioResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RepairApRadioResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *RepairApRadioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RepairApRadioResponseBody) SetData(v map[string]interface{}) *RepairApRadioResponseBody {
	s.Data = v
	return s
}

func (s *RepairApRadioResponseBody) SetErrorCode(v bool) *RepairApRadioResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RepairApRadioResponseBody) SetErrorMessage(v string) *RepairApRadioResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RepairApRadioResponseBody) SetIsSuccess(v bool) *RepairApRadioResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *RepairApRadioResponseBody) SetRequestId(v string) *RepairApRadioResponseBody {
	s.RequestId = &v
	return s
}

func (s *RepairApRadioResponseBody) Validate() error {
	return dara.Validate(s)
}
