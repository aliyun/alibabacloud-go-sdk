// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomerInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v int64) *FindCustomerInfoRequest
	GetUserId() *int64
}

type FindCustomerInfoRequest struct {
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s FindCustomerInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerInfoRequest) GoString() string {
	return s.String()
}

func (s *FindCustomerInfoRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *FindCustomerInfoRequest) SetUserId(v int64) *FindCustomerInfoRequest {
	s.UserId = &v
	return s
}

func (s *FindCustomerInfoRequest) Validate() error {
	return dara.Validate(s)
}
