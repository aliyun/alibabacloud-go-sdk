// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAppFailOverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *InitAppFailOverRequest
	GetApplicationId() *string
}

type InitAppFailOverRequest struct {
	// The application ID.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s InitAppFailOverRequest) String() string {
	return dara.Prettify(s)
}

func (s InitAppFailOverRequest) GoString() string {
	return s.String()
}

func (s *InitAppFailOverRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *InitAppFailOverRequest) SetApplicationId(v string) *InitAppFailOverRequest {
	s.ApplicationId = &v
	return s
}

func (s *InitAppFailOverRequest) Validate() error {
	return dara.Validate(s)
}
