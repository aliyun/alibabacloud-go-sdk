// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaliciousFileWhitelistConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetMaliciousFileWhitelistConfigRequest
	GetConfigId() *int64
}

type GetMaliciousFileWhitelistConfigRequest struct {
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 1
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s GetMaliciousFileWhitelistConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMaliciousFileWhitelistConfigRequest) GoString() string {
	return s.String()
}

func (s *GetMaliciousFileWhitelistConfigRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetMaliciousFileWhitelistConfigRequest) SetConfigId(v int64) *GetMaliciousFileWhitelistConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigRequest) Validate() error {
	return dara.Validate(s)
}
