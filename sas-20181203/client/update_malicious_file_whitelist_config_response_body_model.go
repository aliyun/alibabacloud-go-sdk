// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaliciousFileWhitelistConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMaliciousFileWhitelistConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateMaliciousFileWhitelistConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMaliciousFileWhitelistConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMaliciousFileWhitelistConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMaliciousFileWhitelistConfigResponseBody
	GetSuccess() *bool
}

type UpdateMaliciousFileWhitelistConfigResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-XXXXXXXXX
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

func (s UpdateMaliciousFileWhitelistConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaliciousFileWhitelistConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) SetCode(v string) *UpdateMaliciousFileWhitelistConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) SetHttpStatusCode(v int32) *UpdateMaliciousFileWhitelistConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) SetMessage(v string) *UpdateMaliciousFileWhitelistConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) SetRequestId(v string) *UpdateMaliciousFileWhitelistConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) SetSuccess(v bool) *UpdateMaliciousFileWhitelistConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
