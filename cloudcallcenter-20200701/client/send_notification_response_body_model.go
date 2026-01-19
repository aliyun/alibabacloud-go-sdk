// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendNotificationResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *SendNotificationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SendNotificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendNotificationResponseBody
	GetRequestId() *string
}

type SendNotificationResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39138031-****-****-****-3AEE8F5DAC13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *SendNotificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendNotificationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SendNotificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendNotificationResponseBody) SetCode(v string) *SendNotificationResponseBody {
	s.Code = &v
	return s
}

func (s *SendNotificationResponseBody) SetHttpStatusCode(v string) *SendNotificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendNotificationResponseBody) SetMessage(v string) *SendNotificationResponseBody {
	s.Message = &v
	return s
}

func (s *SendNotificationResponseBody) SetRequestId(v string) *SendNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}
