// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerId(v string) *GetListenerAttributeRequest
	GetListenerId() *string
}

type GetListenerAttributeRequest struct {
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
