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
	SetSuccess(v bool) *CreateImageResponseBody
	GetSuccess() *bool
}

type CreateImageResponseBody struct {
	// The ID of the created image.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID, which is used for locating logs and troubleshooting.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CreateImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponseBody) SetSuccess(v bool) *CreateImageResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageResponseBody) Validate() error {
	return dara.Validate(s)
}
