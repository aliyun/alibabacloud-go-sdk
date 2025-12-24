// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishPlaybookResponseBody
	GetRequestId() *string
}

type PublishPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C513FCEA-D71F-5E50-ADC4-FCF8C5DCF6BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *PublishPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishPlaybookResponseBody) SetRequestId(v string) *PublishPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishPlaybookResponseBody) Validate() error {
	return dara.Validate(s)
}
