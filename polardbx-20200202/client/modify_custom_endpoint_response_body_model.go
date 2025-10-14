// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCustomEndpointResponseBody
	GetRequestId() *string
}

type ModifyCustomEndpointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1A586DCB-39A6-4050-81CC-C7BD4CCDB49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomEndpointResponseBody) SetRequestId(v string) *ModifyCustomEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
