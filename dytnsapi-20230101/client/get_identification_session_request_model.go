// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentificationSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetIdentificationSessionRequest
	GetAuthCode() *string
}

type GetIdentificationSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
}

func (s GetIdentificationSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationSessionRequest) GoString() string {
	return s.String()
}

func (s *GetIdentificationSessionRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetIdentificationSessionRequest) SetAuthCode(v string) *GetIdentificationSessionRequest {
	s.AuthCode = &v
	return s
}

func (s *GetIdentificationSessionRequest) Validate() error {
	return dara.Validate(s)
}
