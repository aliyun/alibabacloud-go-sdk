// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckScopeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCheckScopeConfigRequest
	GetConfigId() *int64
	SetResourceDirectoryAccountId(v int64) *GetCheckScopeConfigRequest
	GetResourceDirectoryAccountId() *int64
}

type GetCheckScopeConfigRequest struct {
	// The ID of the configuration. This parameter is optional. If you do not specify this parameter, a default ID is generated.
	//
	// example:
	//
	// 435f626256ebf564cf5ba966a539****
	ConfigId                   *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s GetCheckScopeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckScopeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetCheckScopeConfigRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCheckScopeConfigRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetCheckScopeConfigRequest) SetConfigId(v int64) *GetCheckScopeConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCheckScopeConfigRequest) SetResourceDirectoryAccountId(v int64) *GetCheckScopeConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetCheckScopeConfigRequest) Validate() error {
	return dara.Validate(s)
}
