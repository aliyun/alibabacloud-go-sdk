// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKeyCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ApiKeyCreateCmd
	GetClientId() *int64
}

type ApiKeyCreateCmd struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s ApiKeyCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyCreateCmd) GoString() string {
	return s.String()
}

func (s *ApiKeyCreateCmd) GetClientId() *int64 {
	return s.ClientId
}

func (s *ApiKeyCreateCmd) SetClientId(v int64) *ApiKeyCreateCmd {
	s.ClientId = &v
	return s
}

func (s *ApiKeyCreateCmd) Validate() error {
	return dara.Validate(s)
}
