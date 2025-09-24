// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishCodeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeSourceId(v string) *PublishCodeSourceResponseBody
	GetCodeSourceId() *string
	SetRequestId(v string) *PublishCodeSourceResponseBody
	GetRequestId() *string
}

type PublishCodeSourceResponseBody struct {
	// The ID of the code source that is successfully published.
	//
	// example:
	//
	// code-a797*******
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishCodeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PublishCodeSourceResponseBody) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *PublishCodeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishCodeSourceResponseBody) SetCodeSourceId(v string) *PublishCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *PublishCodeSourceResponseBody) SetRequestId(v string) *PublishCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishCodeSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
