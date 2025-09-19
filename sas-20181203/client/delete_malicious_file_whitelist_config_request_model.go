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
	// The ID of the alert whitelist rule of sensitive files that are detected by using the agentless detection feature. You can call the [ListMaliciousFileWhitelistConfigs](~~ListMaliciousFileWhitelistConfigs~~) operation to query the IDs of alert whitelist rules.
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
