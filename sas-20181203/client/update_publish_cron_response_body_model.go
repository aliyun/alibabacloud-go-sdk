// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishCronResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublishCronResponseBody
	GetRequestId() *string
}

type UpdatePublishCronResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1052B989-305B-50A5-B5F5-998450******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublishCronResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishCronResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublishCronResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublishCronResponseBody) SetRequestId(v string) *UpdatePublishCronResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublishCronResponseBody) Validate() error {
	return dara.Validate(s)
}
