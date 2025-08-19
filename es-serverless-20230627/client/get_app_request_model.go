// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetailed(v bool) *GetAppRequest
	GetDetailed() *bool
}

type GetAppRequest struct {
	// example:
	//
	// false
	Detailed *bool `json:"detailed,omitempty" xml:"detailed,omitempty"`
}

func (s GetAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) GetDetailed() *bool {
	return s.Detailed
}

func (s *GetAppRequest) SetDetailed(v bool) *GetAppRequest {
	s.Detailed = &v
	return s
}

func (s *GetAppRequest) Validate() error {
	return dara.Validate(s)
}
