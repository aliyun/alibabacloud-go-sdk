// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteShortUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteShortUrlResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteShortUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteShortUrlResponseBody
	GetRequestId() *string
}

type DeleteShortUrlResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
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
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteShortUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteShortUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteShortUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteShortUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteShortUrlResponseBody) SetCode(v string) *DeleteShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteShortUrlResponseBody) SetMessage(v string) *DeleteShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteShortUrlResponseBody) SetRequestId(v string) *DeleteShortUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteShortUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
