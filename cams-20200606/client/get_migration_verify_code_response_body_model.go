// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationVerifyCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetMigrationVerifyCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetMigrationVerifyCodeResponseBody
	GetCode() *string
	SetData(v *GetMigrationVerifyCodeResponseBodyData) *GetMigrationVerifyCodeResponseBody
	GetData() *GetMigrationVerifyCodeResponseBodyData
	SetMessage(v string) *GetMigrationVerifyCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMigrationVerifyCodeResponseBody
	GetRequestId() *string
}

type GetMigrationVerifyCodeResponseBody struct {
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
	Data *GetMigrationVerifyCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
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
}

func (s GetMigrationVerifyCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetMigrationVerifyCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMigrationVerifyCodeResponseBody) GetData() *GetMigrationVerifyCodeResponseBodyData {
	return s.Data
}

func (s *GetMigrationVerifyCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMigrationVerifyCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMigrationVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *GetMigrationVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetCode(v string) *GetMigrationVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetData(v *GetMigrationVerifyCodeResponseBodyData) *GetMigrationVerifyCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetMessage(v string) *GetMigrationVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) SetRequestId(v string) *GetMigrationVerifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMigrationVerifyCodeResponseBodyData struct {
	// The ID of the number.
	//
	// example:
	//
	// 82828893332
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Phone number.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetMigrationVerifyCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationVerifyCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetMigrationVerifyCodeResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetMigrationVerifyCodeResponseBodyData) SetId(v string) *GetMigrationVerifyCodeResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBodyData) SetPhoneNumber(v string) *GetMigrationVerifyCodeResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetMigrationVerifyCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
