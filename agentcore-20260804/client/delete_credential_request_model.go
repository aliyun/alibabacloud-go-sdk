// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCredentialRequest
	GetClientToken() *string
}

type DeleteCredentialRequest struct {
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCredentialRequest) SetClientToken(v string) *DeleteCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCredentialRequest) Validate() error {
	return dara.Validate(s)
}
