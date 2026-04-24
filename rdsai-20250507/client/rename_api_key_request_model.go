// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *RenameApiKeyRequest
	GetApiKey() *string
	SetInstanceId(v string) *RenameApiKeyRequest
	GetInstanceId() *string
	SetKeyName(v string) *RenameApiKeyRequest
	GetKeyName() *string
}

type RenameApiKeyRequest struct {
	// API KEY
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
	// This parameter is required.
	//
	// example:
	//
	// api-*****
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
}

func (s RenameApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameApiKeyRequest) GoString() string {
	return s.String()
}

func (s *RenameApiKeyRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RenameApiKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenameApiKeyRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *RenameApiKeyRequest) SetApiKey(v string) *RenameApiKeyRequest {
	s.ApiKey = &v
	return s
}

func (s *RenameApiKeyRequest) SetInstanceId(v string) *RenameApiKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *RenameApiKeyRequest) SetKeyName(v string) *RenameApiKeyRequest {
	s.KeyName = &v
	return s
}

func (s *RenameApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
