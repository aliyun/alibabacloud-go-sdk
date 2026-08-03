// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GetApiKeyRequest
	GetKeyId() *string
	SetRegionId(v string) *GetApiKeyRequest
	GetRegionId() *string
}

type GetApiKeyRequest struct {
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

func (s GetApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyRequest) GoString() string {
	return s.String()
}

func (s *GetApiKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GetApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApiKeyRequest) SetKeyId(v string) *GetApiKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GetApiKeyRequest) SetRegionId(v string) *GetApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GetApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
