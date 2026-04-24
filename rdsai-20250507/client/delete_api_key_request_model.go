// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DeleteApiKeyRequest
	GetApiKey() *string
	SetInstanceId(v string) *DeleteApiKeyRequest
	GetInstanceId() *string
}

type DeleteApiKeyRequest struct {
	// Api Key
	//
	// This parameter is required.
	//
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeleteApiKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApiKeyRequest) SetApiKey(v string) *DeleteApiKeyRequest {
	s.ApiKey = &v
	return s
}

func (s *DeleteApiKeyRequest) SetInstanceId(v string) *DeleteApiKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
