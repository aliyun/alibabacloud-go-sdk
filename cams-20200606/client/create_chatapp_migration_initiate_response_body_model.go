// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappMigrationInitiateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateChatappMigrationInitiateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateChatappMigrationInitiateResponseBody
	GetCode() *string
	SetData(v *CreateChatappMigrationInitiateResponseBodyData) *CreateChatappMigrationInitiateResponseBody
	GetData() *CreateChatappMigrationInitiateResponseBodyData
	SetMessage(v string) *CreateChatappMigrationInitiateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatappMigrationInitiateResponseBody
	GetRequestId() *string
}

type CreateChatappMigrationInitiateResponseBody struct {
	// The information about the request denial..
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- A value of OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *CreateChatappMigrationInitiateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateChatappMigrationInitiateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappMigrationInitiateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateChatappMigrationInitiateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatappMigrationInitiateResponseBody) GetData() *CreateChatappMigrationInitiateResponseBodyData {
	return s.Data
}

func (s *CreateChatappMigrationInitiateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatappMigrationInitiateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatappMigrationInitiateResponseBody) SetAccessDeniedDetail(v string) *CreateChatappMigrationInitiateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetCode(v string) *CreateChatappMigrationInitiateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetData(v *CreateChatappMigrationInitiateResponseBodyData) *CreateChatappMigrationInitiateResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetMessage(v string) *CreateChatappMigrationInitiateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) SetRequestId(v string) *CreateChatappMigrationInitiateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateChatappMigrationInitiateResponseBodyData struct {
	// The ID of the mobile number.
	//
	// example:
	//
	// 82828893332
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 8613900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The state of the mobile number. Only MIGRATING may be returned, which indicates that the mobile number is being migrated.
	//
	// example:
	//
	// MIGRATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateChatappMigrationInitiateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappMigrationInitiateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateChatappMigrationInitiateResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateChatappMigrationInitiateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateChatappMigrationInitiateResponseBodyData) SetId(v string) *CreateChatappMigrationInitiateResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBodyData) SetPhoneNumber(v string) *CreateChatappMigrationInitiateResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBodyData) SetStatus(v string) *CreateChatappMigrationInitiateResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
