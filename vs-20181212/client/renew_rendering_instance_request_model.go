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
	// Enable or disable auto-renewal. Valid values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The duration of the subscription. Valid values are 1 (default), 2, 3, 4, 5, 6, 7, 8, 9, 12. A value of 12 is converted to one year; other values are in months.
	//
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// Cloud application service instance ID.
	//
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
