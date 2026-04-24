// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateApiKeyQuotaRequest
	GetInstanceId() *string
	SetKeys(v []*UpdateApiKeyQuotaRequestKeys) *UpdateApiKeyQuotaRequest
	GetKeys() []*UpdateApiKeyQuotaRequestKeys
}

type UpdateApiKeyQuotaRequest struct {
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keys       []*UpdateApiKeyQuotaRequestKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s UpdateApiKeyQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApiKeyQuotaRequest) GetKeys() []*UpdateApiKeyQuotaRequestKeys {
	return s.Keys
}

func (s *UpdateApiKeyQuotaRequest) SetInstanceId(v string) *UpdateApiKeyQuotaRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApiKeyQuotaRequest) SetKeys(v []*UpdateApiKeyQuotaRequestKeys) *UpdateApiKeyQuotaRequest {
	s.Keys = v
	return s
}

func (s *UpdateApiKeyQuotaRequest) Validate() error {
	if s.Keys != nil {
		for _, item := range s.Keys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApiKeyQuotaRequestKeys struct {
	// API KEY
	//
	// example:
	//
	// sk-rds-xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 0.2
	LimitRate *float64 `json:"LimitRate,omitempty" xml:"LimitRate,omitempty"`
	// example:
	//
	// fixed
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
	// example:
	//
	// 100000
	TokenQuota *int64 `json:"TokenQuota,omitempty" xml:"TokenQuota,omitempty"`
}

func (s UpdateApiKeyQuotaRequestKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaRequestKeys) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaRequestKeys) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateApiKeyQuotaRequestKeys) GetLimitRate() *float64 {
	return s.LimitRate
}

func (s *UpdateApiKeyQuotaRequestKeys) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateApiKeyQuotaRequestKeys) GetTokenQuota() *int64 {
	return s.TokenQuota
}

func (s *UpdateApiKeyQuotaRequestKeys) SetApiKey(v string) *UpdateApiKeyQuotaRequestKeys {
	s.ApiKey = &v
	return s
}

func (s *UpdateApiKeyQuotaRequestKeys) SetLimitRate(v float64) *UpdateApiKeyQuotaRequestKeys {
	s.LimitRate = &v
	return s
}

func (s *UpdateApiKeyQuotaRequestKeys) SetLimitType(v string) *UpdateApiKeyQuotaRequestKeys {
	s.LimitType = &v
	return s
}

func (s *UpdateApiKeyQuotaRequestKeys) SetTokenQuota(v int64) *UpdateApiKeyQuotaRequestKeys {
	s.TokenQuota = &v
	return s
}

func (s *UpdateApiKeyQuotaRequestKeys) Validate() error {
	return dara.Validate(s)
}
