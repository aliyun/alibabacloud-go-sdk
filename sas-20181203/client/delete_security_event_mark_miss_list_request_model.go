// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityEventMarkMissListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *DeleteSecurityEventMarkMissListRequest
	GetIds() []*int64
	SetResourceDirectoryAccountId(v int64) *DeleteSecurityEventMarkMissListRequest
	GetResourceDirectoryAccountId() *int64
	SetResourceOwnerId(v int64) *DeleteSecurityEventMarkMissListRequest
	GetResourceOwnerId() *int64
}

type DeleteSecurityEventMarkMissListRequest struct {
	// The IDs of custom defense rule.
	Ids                        []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	ResourceDirectoryAccountId *int64   `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	ResourceOwnerId            *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteSecurityEventMarkMissListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityEventMarkMissListRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityEventMarkMissListRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteSecurityEventMarkMissListRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DeleteSecurityEventMarkMissListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSecurityEventMarkMissListRequest) SetIds(v []*int64) *DeleteSecurityEventMarkMissListRequest {
	s.Ids = v
	return s
}

func (s *DeleteSecurityEventMarkMissListRequest) SetResourceDirectoryAccountId(v int64) *DeleteSecurityEventMarkMissListRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DeleteSecurityEventMarkMissListRequest) SetResourceOwnerId(v int64) *DeleteSecurityEventMarkMissListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSecurityEventMarkMissListRequest) Validate() error {
	return dara.Validate(s)
}
