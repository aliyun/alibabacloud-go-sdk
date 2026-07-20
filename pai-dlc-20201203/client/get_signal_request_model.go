// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSignalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *GetSignalRequest
	GetToken() *string
}

type GetSignalRequest struct {
	// The temporary token used for authentication.
	//
	// example:
	//
	// eyXXXX-XXXX.XXXXX
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetSignalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSignalRequest) GoString() string {
	return s.String()
}

func (s *GetSignalRequest) GetToken() *string {
	return s.Token
}

func (s *GetSignalRequest) SetToken(v string) *GetSignalRequest {
	s.Token = &v
	return s
}

func (s *GetSignalRequest) Validate() error {
	return dara.Validate(s)
}
