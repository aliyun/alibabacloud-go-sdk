// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishDatasetResponseBody
	GetRequestId() *string
}

type PublishDatasetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishDatasetResponseBody) SetRequestId(v string) *PublishDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
