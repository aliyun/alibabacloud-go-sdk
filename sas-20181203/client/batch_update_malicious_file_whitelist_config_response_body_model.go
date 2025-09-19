// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateMaliciousFileWhitelistConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchUpdateMaliciousFileWhitelistConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BatchUpdateMaliciousFileWhitelistConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchUpdateMaliciousFileWhitelistConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchUpdateMaliciousFileWhitelistConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchUpdateMaliciousFileWhitelistConfigResponseBody
	GetSuccess() *bool
}

type BatchUpdateMaliciousFileWhitelistConfigResponseBody struct {
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
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
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

func (s BatchUpdateMaliciousFileWhitelistConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateMaliciousFileWhitelistConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) SetCode(v string) *BatchUpdateMaliciousFileWhitelistConfigResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) SetHttpStatusCode(v int32) *BatchUpdateMaliciousFileWhitelistConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) SetMessage(v string) *BatchUpdateMaliciousFileWhitelistConfigResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) SetRequestId(v string) *BatchUpdateMaliciousFileWhitelistConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) SetSuccess(v bool) *BatchUpdateMaliciousFileWhitelistConfigResponseBody {
	s.Success = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
