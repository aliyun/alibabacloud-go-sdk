// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialSysomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckOnly(v bool) *InitialSysomRequest
	GetCheckOnly() *bool
	SetSource(v string) *InitialSysomRequest
	GetSource() *string
}

type InitialSysomRequest struct {
	CheckOnly *bool   `json:"check_only,omitempty" xml:"check_only,omitempty"`
	Source    *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s InitialSysomRequest) String() string {
	return dara.Prettify(s)
}

func (s InitialSysomRequest) GoString() string {
	return s.String()
}

func (s *InitialSysomRequest) GetCheckOnly() *bool {
	return s.CheckOnly
}

func (s *InitialSysomRequest) GetSource() *string {
	return s.Source
}

func (s *InitialSysomRequest) SetCheckOnly(v bool) *InitialSysomRequest {
	s.CheckOnly = &v
	return s
}

func (s *InitialSysomRequest) SetSource(v string) *InitialSysomRequest {
	s.Source = &v
	return s
}

func (s *InitialSysomRequest) Validate() error {
	return dara.Validate(s)
}
