// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *PublishImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *PublishImageResponseBody
	GetRequestId() *string
}

type PublishImageResponseBody struct {
	// The image ID.
	//
	// example:
	//
	// image-dk******fa
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishImageResponseBody) GoString() string {
	return s.String()
}

func (s *PublishImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *PublishImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishImageResponseBody) SetImageId(v string) *PublishImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *PublishImageResponseBody) SetRequestId(v string) *PublishImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishImageResponseBody) Validate() error {
	return dara.Validate(s)
}
