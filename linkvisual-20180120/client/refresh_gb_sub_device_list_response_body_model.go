// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshGbSubDeviceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefreshGbSubDeviceListResponseBody
	GetCode() *string
	SetErrorMessage(v string) *RefreshGbSubDeviceListResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RefreshGbSubDeviceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RefreshGbSubDeviceListResponseBody
	GetSuccess() *bool
}

type RefreshGbSubDeviceListResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshGbSubDeviceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshGbSubDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshGbSubDeviceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *RefreshGbSubDeviceListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RefreshGbSubDeviceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshGbSubDeviceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefreshGbSubDeviceListResponseBody) SetCode(v string) *RefreshGbSubDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) SetErrorMessage(v string) *RefreshGbSubDeviceListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) SetRequestId(v string) *RefreshGbSubDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) SetSuccess(v bool) *RefreshGbSubDeviceListResponseBody {
	s.Success = &v
	return s
}

func (s *RefreshGbSubDeviceListResponseBody) Validate() error {
	return dara.Validate(s)
}
