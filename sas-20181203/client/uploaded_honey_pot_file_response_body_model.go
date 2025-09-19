// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadedHoneyPotFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadedHoneyPotFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UploadedHoneyPotFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UploadedHoneyPotFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadedHoneyPotFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadedHoneyPotFileResponseBody
	GetSuccess() *bool
}

type UploadedHoneyPotFileResponseBody struct {
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
	// 427F89F8-6DFE-57CC-9593-3487CA93****
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

func (s UploadedHoneyPotFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadedHoneyPotFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadedHoneyPotFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadedHoneyPotFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UploadedHoneyPotFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadedHoneyPotFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadedHoneyPotFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadedHoneyPotFileResponseBody) SetCode(v string) *UploadedHoneyPotFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadedHoneyPotFileResponseBody) SetHttpStatusCode(v int32) *UploadedHoneyPotFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadedHoneyPotFileResponseBody) SetMessage(v string) *UploadedHoneyPotFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadedHoneyPotFileResponseBody) SetRequestId(v string) *UploadedHoneyPotFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadedHoneyPotFileResponseBody) SetSuccess(v bool) *UploadedHoneyPotFileResponseBody {
	s.Success = &v
	return s
}

func (s *UploadedHoneyPotFileResponseBody) Validate() error {
	return dara.Validate(s)
}
