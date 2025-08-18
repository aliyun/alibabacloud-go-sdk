// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageFromAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CreateImageFromAppInstanceGroupResponseBody
	GetImageId() *string
	SetRequestId(v string) *CreateImageFromAppInstanceGroupResponseBody
	GetRequestId() *string
}

type CreateImageFromAppInstanceGroupResponseBody struct {
	// The image ID.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageFromAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CreateImageFromAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageFromAppInstanceGroupResponseBody) SetImageId(v string) *CreateImageFromAppInstanceGroupResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponseBody) SetRequestId(v string) *CreateImageFromAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
