// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagerInfoByAuthCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetManagerInfoByAuthCodeRequest
	GetAuthCode() *string
}

type GetManagerInfoByAuthCodeRequest struct {
	// The authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// e49cd070452f0044813a467d4743****
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
}

func (s GetManagerInfoByAuthCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagerInfoByAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *GetManagerInfoByAuthCodeRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetManagerInfoByAuthCodeRequest) SetAuthCode(v string) *GetManagerInfoByAuthCodeRequest {
	s.AuthCode = &v
	return s
}

func (s *GetManagerInfoByAuthCodeRequest) Validate() error {
	return dara.Validate(s)
}
