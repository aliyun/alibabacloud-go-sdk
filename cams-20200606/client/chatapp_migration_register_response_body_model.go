// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappMigrationRegisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappMigrationRegisterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappMigrationRegisterResponseBody
	GetCode() *string
	SetMessage(v string) *ChatappMigrationRegisterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappMigrationRegisterResponseBody
	GetRequestId() *string
}

type ChatappMigrationRegisterResponseBody struct {
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
	// The error message.
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

func (s ChatappMigrationRegisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappMigrationRegisterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappMigrationRegisterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappMigrationRegisterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappMigrationRegisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappMigrationRegisterResponseBody) SetAccessDeniedDetail(v string) *ChatappMigrationRegisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) SetCode(v string) *ChatappMigrationRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) SetMessage(v string) *ChatappMigrationRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) SetRequestId(v string) *ChatappMigrationRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappMigrationRegisterResponseBody) Validate() error {
	return dara.Validate(s)
}
