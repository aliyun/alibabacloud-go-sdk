// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMpk(v string) *ReleaseAgAccountRequest
	GetMpk() *string
	SetPk(v string) *ReleaseAgAccountRequest
	GetPk() *string
	SetReleaseReason(v string) *ReleaseAgAccountRequest
	GetReleaseReason() *string
}

type ReleaseAgAccountRequest struct {
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk            *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
}

func (s ReleaseAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAgAccountRequest) GoString() string {
	return s.String()
}

func (s *ReleaseAgAccountRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ReleaseAgAccountRequest) GetPk() *string {
	return s.Pk
}

func (s *ReleaseAgAccountRequest) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *ReleaseAgAccountRequest) SetMpk(v string) *ReleaseAgAccountRequest {
	s.Mpk = &v
	return s
}

func (s *ReleaseAgAccountRequest) SetPk(v string) *ReleaseAgAccountRequest {
	s.Pk = &v
	return s
}

func (s *ReleaseAgAccountRequest) SetReleaseReason(v string) *ReleaseAgAccountRequest {
	s.ReleaseReason = &v
	return s
}

func (s *ReleaseAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
