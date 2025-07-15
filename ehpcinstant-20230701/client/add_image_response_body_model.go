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
	SetSuccess(v bool) *AddImageResponseBody
	GetSuccess() *bool
}

type AddImageResponseBody struct {
	// example:
	//
	// m-bp1akkkr1rkxtb******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *AddImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddImageResponseBody) SetImageId(v string) *AddImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

func (s *AddImageResponseBody) Validate() error {
	return dara.Validate(s)
}
