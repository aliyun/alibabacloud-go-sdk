// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaliciousFileWhitelistConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteMaliciousFileWhitelistConfigRequest
	GetConfigId() *int64
}

type DeleteMaliciousFileWhitelistConfigRequest struct {
	// The ID of the whitelist rule for agentless detection of sensitive file alerts. You can call [ListMaliciousFileWhitelistConfigs](~~ListMaliciousFileWhitelistConfigs~~) to obtain the ID.
	//
	// example:
	//
	// 1
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s DeleteMaliciousFileWhitelistConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaliciousFileWhitelistConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaliciousFileWhitelistConfigRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteMaliciousFileWhitelistConfigRequest) SetConfigId(v int64) *DeleteMaliciousFileWhitelistConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigRequest) Validate() error {
	return dara.Validate(s)
}
