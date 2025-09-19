// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMaliciousFileWhitelistConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchDeleteMaliciousFileWhitelistConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BatchDeleteMaliciousFileWhitelistConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchDeleteMaliciousFileWhitelistConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteMaliciousFileWhitelistConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDeleteMaliciousFileWhitelistConfigResponseBody
	GetSuccess() *bool
}

type BatchDeleteMaliciousFileWhitelistConfigResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The status code 200 indicates that the request was successful.
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
	// A3D7C47D-3F11-57BB-90E8-E5C20C61****
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

func (s BatchDeleteMaliciousFileWhitelistConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMaliciousFileWhitelistConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) SetCode(v string) *BatchDeleteMaliciousFileWhitelistConfigResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) SetHttpStatusCode(v int32) *BatchDeleteMaliciousFileWhitelistConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) SetMessage(v string) *BatchDeleteMaliciousFileWhitelistConfigResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) SetRequestId(v string) *BatchDeleteMaliciousFileWhitelistConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) SetSuccess(v bool) *BatchDeleteMaliciousFileWhitelistConfigResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
