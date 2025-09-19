// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStrictEventNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStrictEventNameResponseBody
	GetRequestId() *string
}

type UpdateStrictEventNameResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStrictEventNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStrictEventNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStrictEventNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStrictEventNameResponseBody) SetRequestId(v string) *UpdateStrictEventNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStrictEventNameResponseBody) Validate() error {
	return dara.Validate(s)
}
