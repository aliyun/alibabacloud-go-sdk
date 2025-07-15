// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CreateImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *CreateImageResponseBody
	GetRequestId() *string
}

type CreateImageResponseBody struct {
	// The ID of the image.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CreateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponseBody) Validate() error {
	return dara.Validate(s)
}
