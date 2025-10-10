// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenEmasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenEmasServiceRequest
	GetOwnerId() *int64
}

type OpenEmasServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenEmasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenEmasServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenEmasServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenEmasServiceRequest) SetOwnerId(v int64) *OpenEmasServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenEmasServiceRequest) Validate() error {
	return dara.Validate(s)
}
