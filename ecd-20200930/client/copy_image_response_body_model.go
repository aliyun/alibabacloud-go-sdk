// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CopyImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *CopyImageResponseBody
	GetRequestId() *string
}

type CopyImageResponseBody struct {
	// The ID of the image that is being copied.
	//
	// example:
	//
	// m-2g65ljy3ynrdq****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 28A40F12-F340-442B-A35F-46EF6A03****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CopyImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CopyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyImageResponseBody) SetImageId(v string) *CopyImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CopyImageResponseBody) SetRequestId(v string) *CopyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyImageResponseBody) Validate() error {
	return dara.Validate(s)
}
