// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerInformationResponseBody
	GetCode() *string
	SetData(v *GetCustomerInformationResponseBodyData) *GetCustomerInformationResponseBody
	GetData() *GetCustomerInformationResponseBodyData
	SetMessage(v string) *GetCustomerInformationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerInformationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerInformationResponseBody
	GetSuccess() *bool
}

type GetCustomerInformationResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCustomerInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerInformationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerInformationResponseBody) GetData() *GetCustomerInformationResponseBodyData {
	return s.Data
}

func (s *GetCustomerInformationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerInformationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerInformationResponseBody) SetCode(v string) *GetCustomerInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerInformationResponseBody) SetData(v *GetCustomerInformationResponseBodyData) *GetCustomerInformationResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerInformationResponseBody) SetMessage(v string) *GetCustomerInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerInformationResponseBody) SetRequestId(v string) *GetCustomerInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerInformationResponseBody) SetSuccess(v bool) *GetCustomerInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerInformationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerInformationResponseBodyData struct {
	Biz              *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	CustomerCategory *string `json:"CustomerCategory,omitempty" xml:"CustomerCategory,omitempty"`
	Website          *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s GetCustomerInformationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerInformationResponseBodyData) GetBiz() *string {
	return s.Biz
}

func (s *GetCustomerInformationResponseBodyData) GetCustomerCategory() *string {
	return s.CustomerCategory
}

func (s *GetCustomerInformationResponseBodyData) GetWebsite() *string {
	return s.Website
}

func (s *GetCustomerInformationResponseBodyData) SetBiz(v string) *GetCustomerInformationResponseBodyData {
	s.Biz = &v
	return s
}

func (s *GetCustomerInformationResponseBodyData) SetCustomerCategory(v string) *GetCustomerInformationResponseBodyData {
	s.CustomerCategory = &v
	return s
}

func (s *GetCustomerInformationResponseBodyData) SetWebsite(v string) *GetCustomerInformationResponseBodyData {
	s.Website = &v
	return s
}

func (s *GetCustomerInformationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
