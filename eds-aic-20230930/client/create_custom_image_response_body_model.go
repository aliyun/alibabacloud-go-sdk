// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CreateCustomImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *CreateCustomImageResponseBody
	GetRequestId() *string
}

type CreateCustomImageResponseBody struct {
	// The ID of the custom image.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 20393E53-8FF1-524C-B494-B478A5369733
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CreateCustomImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomImageResponseBody) SetImageId(v string) *CreateCustomImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateCustomImageResponseBody) SetRequestId(v string) *CreateCustomImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomImageResponseBody) Validate() error {
	return dara.Validate(s)
}
