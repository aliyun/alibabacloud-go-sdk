// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ResetApiKeyRequest
	GetApiKey() *string
	SetInstanceId(v string) *ResetApiKeyRequest
	GetInstanceId() *string
}

type ResetApiKeyRequest struct {
	// Api Key
	//
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResetApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ResetApiKeyRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *ResetApiKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetApiKeyRequest) SetApiKey(v string) *ResetApiKeyRequest {
	s.ApiKey = &v
	return s
}

func (s *ResetApiKeyRequest) SetInstanceId(v string) *ResetApiKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
