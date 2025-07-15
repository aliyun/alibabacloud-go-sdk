// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UploadImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *UploadImageResponseBody
	GetRequestId() *string
}

type UploadImageResponseBody struct {
	// The ID of the custom image.
	//
	// example:
	//
	// m-d4dwr5tgrgvd****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2CC66B0A-BA3B-5D87-BFBE-11AAAD7A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadImageResponseBody) GoString() string {
	return s.String()
}

func (s *UploadImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *UploadImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadImageResponseBody) SetImageId(v string) *UploadImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *UploadImageResponseBody) SetRequestId(v string) *UploadImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadImageResponseBody) Validate() error {
	return dara.Validate(s)
}
