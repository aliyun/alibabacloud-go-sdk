// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPrivateLinkServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenPrivateLinkServiceRequest
	GetOwnerId() *int64
}

type OpenPrivateLinkServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenPrivateLinkServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenPrivateLinkServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenPrivateLinkServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenPrivateLinkServiceRequest) SetOwnerId(v int64) *OpenPrivateLinkServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenPrivateLinkServiceRequest) Validate() error {
	return dara.Validate(s)
}
