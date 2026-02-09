// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsFirstBbsTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IsFirstBbsTicketResponseBody
	GetCode() *string
	SetData(v bool) *IsFirstBbsTicketResponseBody
	GetData() *bool
	SetMessage(v string) *IsFirstBbsTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *IsFirstBbsTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IsFirstBbsTicketResponseBody
	GetSuccess() *bool
}

type IsFirstBbsTicketResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IsFirstBbsTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsFirstBbsTicketResponseBody) GoString() string {
	return s.String()
}

func (s *IsFirstBbsTicketResponseBody) GetCode() *string {
	return s.Code
}

func (s *IsFirstBbsTicketResponseBody) GetData() *bool {
	return s.Data
}

func (s *IsFirstBbsTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IsFirstBbsTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsFirstBbsTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IsFirstBbsTicketResponseBody) SetCode(v string) *IsFirstBbsTicketResponseBody {
	s.Code = &v
	return s
}

func (s *IsFirstBbsTicketResponseBody) SetData(v bool) *IsFirstBbsTicketResponseBody {
	s.Data = &v
	return s
}

func (s *IsFirstBbsTicketResponseBody) SetMessage(v string) *IsFirstBbsTicketResponseBody {
	s.Message = &v
	return s
}

func (s *IsFirstBbsTicketResponseBody) SetRequestId(v string) *IsFirstBbsTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsFirstBbsTicketResponseBody) SetSuccess(v bool) *IsFirstBbsTicketResponseBody {
	s.Success = &v
	return s
}

func (s *IsFirstBbsTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
