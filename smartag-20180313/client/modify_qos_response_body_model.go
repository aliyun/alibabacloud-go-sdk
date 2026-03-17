// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQosResponseBody
	GetRequestId() *string
}

type ModifyQosResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EE837E9F-BD50-4C2B-9E47-260F9D848480
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQosResponseBody) SetRequestId(v string) *ModifyQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQosResponseBody) Validate() error {
	return dara.Validate(s)
}
