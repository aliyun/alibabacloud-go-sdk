// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupPortalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *SaveApgroupPortalConfigResponseBody
	GetData() []*int64
	SetErrorCode(v int32) *SaveApgroupPortalConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApgroupPortalConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApgroupPortalConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SaveApgroupPortalConfigResponseBody
	GetRequestId() *string
}

type SaveApgroupPortalConfigResponseBody struct {
	Data         []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveApgroupPortalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupPortalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApgroupPortalConfigResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *SaveApgroupPortalConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApgroupPortalConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApgroupPortalConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApgroupPortalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveApgroupPortalConfigResponseBody) SetData(v []*int64) *SaveApgroupPortalConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApgroupPortalConfigResponseBody) SetErrorCode(v int32) *SaveApgroupPortalConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApgroupPortalConfigResponseBody) SetErrorMessage(v string) *SaveApgroupPortalConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApgroupPortalConfigResponseBody) SetIsSuccess(v bool) *SaveApgroupPortalConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApgroupPortalConfigResponseBody) SetRequestId(v string) *SaveApgroupPortalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApgroupPortalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
