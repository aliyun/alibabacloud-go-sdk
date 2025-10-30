// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSecretNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseSecretNoResponseBody
	GetCode() *string
	SetMessage(v string) *ReleaseSecretNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseSecretNoResponseBody
	GetRequestId() *string
}

type ReleaseSecretNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 986BCB6D-C9BF-42F9-91CE-3A990121232
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseSecretNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSecretNoResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSecretNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseSecretNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseSecretNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseSecretNoResponseBody) SetCode(v string) *ReleaseSecretNoResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) SetMessage(v string) *ReleaseSecretNoResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) SetRequestId(v string) *ReleaseSecretNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseSecretNoResponseBody) Validate() error {
	return dara.Validate(s)
}
