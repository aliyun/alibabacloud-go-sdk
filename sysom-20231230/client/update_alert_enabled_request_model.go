// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateAlertEnabledRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateAlertEnabledRequest
	GetId() *int64
}

type UpdateAlertEnabledRequest struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateAlertEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertEnabledRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertEnabledRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAlertEnabledRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAlertEnabledRequest) SetEnabled(v bool) *UpdateAlertEnabledRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateAlertEnabledRequest) SetId(v int64) *UpdateAlertEnabledRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertEnabledRequest) Validate() error {
	return dara.Validate(s)
}
