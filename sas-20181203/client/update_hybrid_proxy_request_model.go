// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHybridProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProxyUuid(v string) *UpdateHybridProxyRequest
	GetProxyUuid() *string
}

type UpdateHybridProxyRequest struct {
	// The UUID of the Security Center agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-proxy-14bbbb37-c4b9-4e86-83bd-137a940a6ec4
	ProxyUuid *string `json:"ProxyUuid,omitempty" xml:"ProxyUuid,omitempty"`
}

func (s UpdateHybridProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHybridProxyRequest) GoString() string {
	return s.String()
}

func (s *UpdateHybridProxyRequest) GetProxyUuid() *string {
	return s.ProxyUuid
}

func (s *UpdateHybridProxyRequest) SetProxyUuid(v string) *UpdateHybridProxyRequest {
	s.ProxyUuid = &v
	return s
}

func (s *UpdateHybridProxyRequest) Validate() error {
	return dara.Validate(s)
}
