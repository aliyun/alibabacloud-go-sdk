// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySmsCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifySmsCodeResponseBody
	GetCode() *string
	SetData(v bool) *VerifySmsCodeResponseBody
	GetData() *bool
	SetMessage(v string) *VerifySmsCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifySmsCodeResponseBody
	GetRequestId() *string
}

type VerifySmsCodeResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifySmsCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifySmsCodeResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifySmsCodeResponseBody) GetData() *bool {
	return s.Data
}

func (s *VerifySmsCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifySmsCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifySmsCodeResponseBody) SetCode(v string) *VerifySmsCodeResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetData(v bool) *VerifySmsCodeResponseBody {
	s.Data = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetMessage(v string) *VerifySmsCodeResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetRequestId(v string) *VerifySmsCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySmsCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
