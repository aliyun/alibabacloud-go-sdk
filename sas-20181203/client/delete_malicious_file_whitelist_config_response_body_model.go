// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaliciousFileWhitelistConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMaliciousFileWhitelistConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMaliciousFileWhitelistConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMaliciousFileWhitelistConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMaliciousFileWhitelistConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMaliciousFileWhitelistConfigResponseBody
	GetSuccess() *bool
}

type DeleteMaliciousFileWhitelistConfigResponseBody struct {
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
	// D03DD0FD-6041-5107-AC00-383E28F1****
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

func (s DeleteMaliciousFileWhitelistConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaliciousFileWhitelistConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) SetCode(v string) *DeleteMaliciousFileWhitelistConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) SetHttpStatusCode(v int32) *DeleteMaliciousFileWhitelistConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) SetMessage(v string) *DeleteMaliciousFileWhitelistConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) SetRequestId(v string) *DeleteMaliciousFileWhitelistConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) SetSuccess(v bool) *DeleteMaliciousFileWhitelistConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
