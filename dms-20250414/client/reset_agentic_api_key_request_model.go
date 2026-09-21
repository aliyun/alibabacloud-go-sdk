// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgenticApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireAfterSeconds(v int32) *ResetAgenticApiKeyRequest
	GetExpireAfterSeconds() *int32
	SetId(v int64) *ResetAgenticApiKeyRequest
	GetId() *int64
}

type ResetAgenticApiKeyRequest struct {
	// The validity period of the new Access Token starting from the time of this reset, in seconds. Valid values: 1 to 31536000 (approximately 365 days). If you do not specify this parameter, the original expiration time of the Access Token is retained. This parameter is required when the target Access Token has already expired. Otherwise, the system retains the past expiration time and issues an Access Token that is invalid upon creation, and the request is rejected.
	//
	// example:
	//
	// 2592000
	ExpireAfterSeconds *int32 `json:"ExpireAfterSeconds,omitempty" xml:"ExpireAfterSeconds,omitempty"`
	// The ID of the data gateway Access Token to reset. This value is the same as the Id returned by the create and query operations. Only the creator of the Access Token can reset it, and the target Access Token cannot be in a revoked state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetAgenticApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ResetAgenticApiKeyRequest) GetExpireAfterSeconds() *int32 {
	return s.ExpireAfterSeconds
}

func (s *ResetAgenticApiKeyRequest) GetId() *int64 {
	return s.Id
}

func (s *ResetAgenticApiKeyRequest) SetExpireAfterSeconds(v int32) *ResetAgenticApiKeyRequest {
	s.ExpireAfterSeconds = &v
	return s
}

func (s *ResetAgenticApiKeyRequest) SetId(v int64) *ResetAgenticApiKeyRequest {
	s.Id = &v
	return s
}

func (s *ResetAgenticApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
