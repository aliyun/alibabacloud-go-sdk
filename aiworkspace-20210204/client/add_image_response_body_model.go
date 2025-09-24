// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *AddImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *AddImageResponseBody
	GetRequestId() *string
}

type AddImageResponseBody struct {
	// The image ID.
	//
	// example:
	//
	// image-4c62******53uor
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *AddImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageResponseBody) SetImageId(v string) *AddImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) Validate() error {
	return dara.Validate(s)
}
