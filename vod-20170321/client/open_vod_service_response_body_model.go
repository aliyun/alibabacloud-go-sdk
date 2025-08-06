// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVodServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenVodServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenVodServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenVodServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OpenVodServiceResponseBody
	GetSuccess() *bool
}

type OpenVodServiceResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenVodServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenVodServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVodServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenVodServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenVodServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenVodServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OpenVodServiceResponseBody) SetCode(v string) *OpenVodServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenVodServiceResponseBody) SetMessage(v string) *OpenVodServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenVodServiceResponseBody) SetRequestId(v string) *OpenVodServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenVodServiceResponseBody) SetSuccess(v bool) *OpenVodServiceResponseBody {
	s.Success = &v
	return s
}

func (s *OpenVodServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
