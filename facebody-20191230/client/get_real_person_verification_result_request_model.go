// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealPersonVerificationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerificationToken(v string) *GetRealPersonVerificationResultRequest
	GetVerificationToken() *string
}

type GetRealPersonVerificationResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cwek23dw24geor89230hf2rw
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GetRealPersonVerificationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealPersonVerificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultRequest) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *GetRealPersonVerificationResultRequest) SetVerificationToken(v string) *GetRealPersonVerificationResultRequest {
	s.VerificationToken = &v
	return s
}

func (s *GetRealPersonVerificationResultRequest) Validate() error {
	return dara.Validate(s)
}
