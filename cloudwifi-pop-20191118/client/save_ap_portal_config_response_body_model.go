// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApPortalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *SaveApPortalConfigResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *SaveApPortalConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApPortalConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApPortalConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SaveApPortalConfigResponseBody
	GetRequestId() *string
}

type SaveApPortalConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveApPortalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApPortalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApPortalConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SaveApPortalConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApPortalConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApPortalConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApPortalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveApPortalConfigResponseBody) SetData(v map[string]interface{}) *SaveApPortalConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApPortalConfigResponseBody) SetErrorCode(v int32) *SaveApPortalConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApPortalConfigResponseBody) SetErrorMessage(v string) *SaveApPortalConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApPortalConfigResponseBody) SetIsSuccess(v bool) *SaveApPortalConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApPortalConfigResponseBody) SetRequestId(v string) *SaveApPortalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApPortalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
