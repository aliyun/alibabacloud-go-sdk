// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappBindWabaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappBindWabaResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappBindWabaResponseBody
	GetCode() *string
	SetData(v *ChatappBindWabaResponseBodyData) *ChatappBindWabaResponseBody
	GetData() *ChatappBindWabaResponseBodyData
	SetMessage(v string) *ChatappBindWabaResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappBindWabaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatappBindWabaResponseBody
	GetSuccess() *bool
}

type ChatappBindWabaResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ChatappBindWabaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatappBindWabaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappBindWabaResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappBindWabaResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappBindWabaResponseBody) GetData() *ChatappBindWabaResponseBodyData {
	return s.Data
}

func (s *ChatappBindWabaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappBindWabaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappBindWabaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatappBindWabaResponseBody) SetAccessDeniedDetail(v string) *ChatappBindWabaResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetCode(v string) *ChatappBindWabaResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetData(v *ChatappBindWabaResponseBodyData) *ChatappBindWabaResponseBody {
	s.Data = v
	return s
}

func (s *ChatappBindWabaResponseBody) SetMessage(v string) *ChatappBindWabaResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetRequestId(v string) *ChatappBindWabaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappBindWabaResponseBody) SetSuccess(v bool) *ChatappBindWabaResponseBody {
	s.Success = &v
	return s
}

func (s *ChatappBindWabaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatappBindWabaResponseBodyData struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// C02029392939939
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ID of the WhatsApp Business Account (WABA).
	//
	// example:
	//
	// 2939828282
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s ChatappBindWabaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChatappBindWabaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaResponseBodyData) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappBindWabaResponseBodyData) GetWabaId() *string {
	return s.WabaId
}

func (s *ChatappBindWabaResponseBodyData) SetCustSpaceId(v string) *ChatappBindWabaResponseBodyData {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappBindWabaResponseBodyData) SetWabaId(v string) *ChatappBindWabaResponseBodyData {
	s.WabaId = &v
	return s
}

func (s *ChatappBindWabaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
