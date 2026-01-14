// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteListenerRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteListenerRequest
	GetClientToken() *string
	SetListenerId(v string) *DeleteListenerRequest
	GetListenerId() *string
}

type DeleteListenerRequest struct {
	// The ID of the GA instance for which you want to delete a listener.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the listener that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DeleteListenerRequest) SetAcceleratorId(v string) *DeleteListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteListenerRequest) Validate() error {
	return dara.Validate(s)
}
