// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateImageResponseBody
	GetCode() *int32
	SetImageId(v string) *CreateImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *CreateImageResponseBody
	GetRequestId() *string
}

type CreateImageResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-5xxgg
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8BEF0D72-9901-5D43-B7D3-8B42AC26C516
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CreateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageResponseBody) SetCode(v int32) *CreateImageResponseBody {
	s.Code = &v
	return s
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
