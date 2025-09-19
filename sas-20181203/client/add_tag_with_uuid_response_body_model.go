// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagWithUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTagWithUuidResponseBody
	GetRequestId() *string
}

type AddTagWithUuidResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagWithUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTagWithUuidResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagWithUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTagWithUuidResponseBody) SetRequestId(v string) *AddTagWithUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagWithUuidResponseBody) Validate() error {
	return dara.Validate(s)
}
