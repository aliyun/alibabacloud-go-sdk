// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaPublishStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMediaPublishStateResponseBody
	GetRequestId() *string
}

type UpdateMediaPublishStateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 91B6CAB9-034C-4E4E-A40B-E7F5C81E688C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaPublishStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaPublishStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaPublishStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaPublishStateResponseBody) SetRequestId(v string) *UpdateMediaPublishStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaPublishStateResponseBody) Validate() error {
	return dara.Validate(s)
}
