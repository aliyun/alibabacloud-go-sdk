// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomerInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindCustomerInfoResponseBody
	GetCode() *string
	SetData(v *FindCustomerInfoResponseBodyData) *FindCustomerInfoResponseBody
	GetData() *FindCustomerInfoResponseBodyData
	SetMessage(v string) *FindCustomerInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindCustomerInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindCustomerInfoResponseBody
	GetSuccess() *bool
}

type FindCustomerInfoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FindCustomerInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindCustomerInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *FindCustomerInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindCustomerInfoResponseBody) GetData() *FindCustomerInfoResponseBodyData {
	return s.Data
}

func (s *FindCustomerInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindCustomerInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindCustomerInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindCustomerInfoResponseBody) SetCode(v string) *FindCustomerInfoResponseBody {
	s.Code = &v
	return s
}

func (s *FindCustomerInfoResponseBody) SetData(v *FindCustomerInfoResponseBodyData) *FindCustomerInfoResponseBody {
	s.Data = v
	return s
}

func (s *FindCustomerInfoResponseBody) SetMessage(v string) *FindCustomerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *FindCustomerInfoResponseBody) SetRequestId(v string) *FindCustomerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindCustomerInfoResponseBody) SetSuccess(v bool) *FindCustomerInfoResponseBody {
	s.Success = &v
	return s
}

func (s *FindCustomerInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindCustomerInfoResponseBodyData struct {
	Biz     *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Website *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s FindCustomerInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindCustomerInfoResponseBodyData) GetBiz() *string {
	return s.Biz
}

func (s *FindCustomerInfoResponseBodyData) GetWebsite() *string {
	return s.Website
}

func (s *FindCustomerInfoResponseBodyData) SetBiz(v string) *FindCustomerInfoResponseBodyData {
	s.Biz = &v
	return s
}

func (s *FindCustomerInfoResponseBodyData) SetWebsite(v string) *FindCustomerInfoResponseBodyData {
	s.Website = &v
	return s
}

func (s *FindCustomerInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
