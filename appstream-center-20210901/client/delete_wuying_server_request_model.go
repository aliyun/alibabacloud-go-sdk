// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWuyingServerId(v string) *DeleteWuyingServerRequest
	GetWuyingServerId() *string
}

type DeleteWuyingServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s DeleteWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteWuyingServerRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *DeleteWuyingServerRequest) SetWuyingServerId(v string) *DeleteWuyingServerRequest {
	s.WuyingServerId = &v
	return s
}

func (s *DeleteWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
