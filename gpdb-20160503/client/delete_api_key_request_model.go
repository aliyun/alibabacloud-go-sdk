// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DeleteApiKeyRequest
	GetKeyId() *string
	SetRegionId(v string) *DeleteApiKeyRequest
	GetRegionId() *string
}

type DeleteApiKeyRequest struct {
	// API KEY ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// api-xxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *DeleteApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApiKeyRequest) SetKeyId(v string) *DeleteApiKeyRequest {
	s.KeyId = &v
	return s
}

func (s *DeleteApiKeyRequest) SetRegionId(v string) *DeleteApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
