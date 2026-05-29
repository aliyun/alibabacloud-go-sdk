// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutomaticWriteOffRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutomaticWriteOffAmount(v int64) *AutomaticWriteOffRequest
	GetAutomaticWriteOffAmount() *int64
	SetAutomaticWriteOffEnabled(v bool) *AutomaticWriteOffRequest
	GetAutomaticWriteOffEnabled() *bool
	SetCustomerUid(v int64) *AutomaticWriteOffRequest
	GetCustomerUid() *int64
	SetLanguage(v string) *AutomaticWriteOffRequest
	GetLanguage() *string
}

type AutomaticWriteOffRequest struct {
	// example:
	//
	// 100
	AutomaticWriteOffAmount *int64 `json:"AutomaticWriteOffAmount,omitempty" xml:"AutomaticWriteOffAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	AutomaticWriteOffEnabled *bool `json:"AutomaticWriteOffEnabled,omitempty" xml:"AutomaticWriteOffEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s AutomaticWriteOffRequest) String() string {
	return dara.Prettify(s)
}

func (s AutomaticWriteOffRequest) GoString() string {
	return s.String()
}

func (s *AutomaticWriteOffRequest) GetAutomaticWriteOffAmount() *int64 {
	return s.AutomaticWriteOffAmount
}

func (s *AutomaticWriteOffRequest) GetAutomaticWriteOffEnabled() *bool {
	return s.AutomaticWriteOffEnabled
}

func (s *AutomaticWriteOffRequest) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *AutomaticWriteOffRequest) GetLanguage() *string {
	return s.Language
}

func (s *AutomaticWriteOffRequest) SetAutomaticWriteOffAmount(v int64) *AutomaticWriteOffRequest {
	s.AutomaticWriteOffAmount = &v
	return s
}

func (s *AutomaticWriteOffRequest) SetAutomaticWriteOffEnabled(v bool) *AutomaticWriteOffRequest {
	s.AutomaticWriteOffEnabled = &v
	return s
}

func (s *AutomaticWriteOffRequest) SetCustomerUid(v int64) *AutomaticWriteOffRequest {
	s.CustomerUid = &v
	return s
}

func (s *AutomaticWriteOffRequest) SetLanguage(v string) *AutomaticWriteOffRequest {
	s.Language = &v
	return s
}

func (s *AutomaticWriteOffRequest) Validate() error {
	return dara.Validate(s)
}
