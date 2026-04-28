// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetConsumerApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ResetConsumerApiKeyRequest
	GetConsumerId() *string
	SetGwClusterId(v string) *ResetConsumerApiKeyRequest
	GetGwClusterId() *string
	SetRegionId(v string) *ResetConsumerApiKeyRequest
	GetRegionId() *string
}

type ResetConsumerApiKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetConsumerApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetConsumerApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ResetConsumerApiKeyRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ResetConsumerApiKeyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ResetConsumerApiKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetConsumerApiKeyRequest) SetConsumerId(v string) *ResetConsumerApiKeyRequest {
	s.ConsumerId = &v
	return s
}

func (s *ResetConsumerApiKeyRequest) SetGwClusterId(v string) *ResetConsumerApiKeyRequest {
	s.GwClusterId = &v
	return s
}

func (s *ResetConsumerApiKeyRequest) SetRegionId(v string) *ResetConsumerApiKeyRequest {
	s.RegionId = &v
	return s
}

func (s *ResetConsumerApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
