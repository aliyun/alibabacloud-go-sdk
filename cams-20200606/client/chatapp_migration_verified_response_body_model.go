// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappMigrationVerifiedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappMigrationVerifiedResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappMigrationVerifiedResponseBody
	GetCode() *string
	SetData(v *ChatappMigrationVerifiedResponseBodyData) *ChatappMigrationVerifiedResponseBody
	GetData() *ChatappMigrationVerifiedResponseBodyData
	SetMessage(v string) *ChatappMigrationVerifiedResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappMigrationVerifiedResponseBody
	GetRequestId() *string
}

type ChatappMigrationVerifiedResponseBody struct {
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
	Data *ChatappMigrationVerifiedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappMigrationVerifiedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationVerifiedResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappMigrationVerifiedResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappMigrationVerifiedResponseBody) GetData() *ChatappMigrationVerifiedResponseBodyData {
	return s.Data
}

func (s *ChatappMigrationVerifiedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappMigrationVerifiedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappMigrationVerifiedResponseBody) SetAccessDeniedDetail(v string) *ChatappMigrationVerifiedResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetCode(v string) *ChatappMigrationVerifiedResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetData(v *ChatappMigrationVerifiedResponseBodyData) *ChatappMigrationVerifiedResponseBody {
	s.Data = v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetMessage(v string) *ChatappMigrationVerifiedResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) SetRequestId(v string) *ChatappMigrationVerifiedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChatappMigrationVerifiedResponseBodyData struct {
	// The ID of the phone number.
	//
	// example:
	//
	// 82828893332
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ChatappMigrationVerifiedResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationVerifiedResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ChatappMigrationVerifiedResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappMigrationVerifiedResponseBodyData) SetId(v string) *ChatappMigrationVerifiedResponseBodyData {
	s.Id = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBodyData) SetPhoneNumber(v string) *ChatappMigrationVerifiedResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappMigrationVerifiedResponseBodyData) Validate() error {
	return dara.Validate(s)
}
