// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *RenewRenderingInstanceRequest
	GetAutoRenew() *bool
	SetPeriod(v string) *RenewRenderingInstanceRequest
	GetPeriod() *string
	SetRenderingInstanceId(v string) *RenewRenderingInstanceRequest
	GetRenderingInstanceId() *string
}

type RenewRenderingInstanceRequest struct {
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s RenewRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewRenderingInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewRenderingInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *RenewRenderingInstanceRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RenewRenderingInstanceRequest) SetAutoRenew(v bool) *RenewRenderingInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewRenderingInstanceRequest) SetPeriod(v string) *RenewRenderingInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewRenderingInstanceRequest) SetRenderingInstanceId(v string) *RenewRenderingInstanceRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *RenewRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
