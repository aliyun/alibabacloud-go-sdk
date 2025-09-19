// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMaliciousFileWhitelistConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigIdList(v []*int64) *BatchDeleteMaliciousFileWhitelistConfigRequest
	GetConfigIdList() []*int64
}

type BatchDeleteMaliciousFileWhitelistConfigRequest struct {
	// The IDs of the whitelist rules. You can call the [ListMaliciousFileWhitelistConfigs](~~ListMaliciousFileWhitelistConfigs~~) operation to query the IDs of whitelist rules.
	ConfigIdList []*int64 `json:"ConfigIdList,omitempty" xml:"ConfigIdList,omitempty" type:"Repeated"`
}

func (s BatchDeleteMaliciousFileWhitelistConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMaliciousFileWhitelistConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteMaliciousFileWhitelistConfigRequest) GetConfigIdList() []*int64 {
	return s.ConfigIdList
}

func (s *BatchDeleteMaliciousFileWhitelistConfigRequest) SetConfigIdList(v []*int64) *BatchDeleteMaliciousFileWhitelistConfigRequest {
	s.ConfigIdList = v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigRequest) Validate() error {
	return dara.Validate(s)
}
