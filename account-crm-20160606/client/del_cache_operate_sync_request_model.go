// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelCacheOperateSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DelCacheOperateSyncRequest
	GetKey() *string
}

type DelCacheOperateSyncRequest struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DelCacheOperateSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s DelCacheOperateSyncRequest) GoString() string {
	return s.String()
}

func (s *DelCacheOperateSyncRequest) GetKey() *string {
	return s.Key
}

func (s *DelCacheOperateSyncRequest) SetKey(v string) *DelCacheOperateSyncRequest {
	s.Key = &v
	return s
}

func (s *DelCacheOperateSyncRequest) Validate() error {
	return dara.Validate(s)
}
