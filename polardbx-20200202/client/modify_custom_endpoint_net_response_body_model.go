// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomEndpointNetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCustomEndpointNetResponseBody
	GetRequestId() *string
}

type ModifyCustomEndpointNetResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomEndpointNetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomEndpointNetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomEndpointNetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomEndpointNetResponseBody) SetRequestId(v string) *ModifyCustomEndpointNetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomEndpointNetResponseBody) Validate() error {
	return dara.Validate(s)
}
