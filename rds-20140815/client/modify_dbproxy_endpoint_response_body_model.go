// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBProxyEndpointResponseBody
	GetRequestId() *string
}

type ModifyDBProxyEndpointResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6B50D92C-1960-4D4F-A290-AFADD6B1A5C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBProxyEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBProxyEndpointResponseBody) SetRequestId(v string) *ModifyDBProxyEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBProxyEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
