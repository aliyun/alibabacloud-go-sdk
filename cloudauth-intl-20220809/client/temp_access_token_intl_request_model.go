// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempAccessTokenIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *TempAccessTokenIntlRequest
	GetType() *string
}

type TempAccessTokenIntlRequest struct {
	// example:
	//
	// none
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TempAccessTokenIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s TempAccessTokenIntlRequest) GoString() string {
	return s.String()
}

func (s *TempAccessTokenIntlRequest) GetType() *string {
	return s.Type
}

func (s *TempAccessTokenIntlRequest) SetType(v string) *TempAccessTokenIntlRequest {
	s.Type = &v
	return s
}

func (s *TempAccessTokenIntlRequest) Validate() error {
	return dara.Validate(s)
}
