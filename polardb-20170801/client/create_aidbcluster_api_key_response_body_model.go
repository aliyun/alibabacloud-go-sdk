// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *CreateAIDBClusterApiKeyResponseBodyApiKey) *CreateAIDBClusterApiKeyResponseBody
	GetApiKey() *CreateAIDBClusterApiKeyResponseBodyApiKey
	SetRequestId(v string) *CreateAIDBClusterApiKeyResponseBody
	GetRequestId() *string
}

type CreateAIDBClusterApiKeyResponseBody struct {
	// The API key.
	ApiKey *CreateAIDBClusterApiKeyResponseBodyApiKey `json:"ApiKey,omitempty" xml:"ApiKey,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAIDBClusterApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterApiKeyResponseBody) GetApiKey() *CreateAIDBClusterApiKeyResponseBodyApiKey {
	return s.ApiKey
}

func (s *CreateAIDBClusterApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIDBClusterApiKeyResponseBody) SetApiKey(v *CreateAIDBClusterApiKeyResponseBodyApiKey) *CreateAIDBClusterApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBody) SetRequestId(v string) *CreateAIDBClusterApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAIDBClusterApiKeyResponseBodyApiKey struct {
	// The API key for model serving.
	//
	// example:
	//
	// xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-12T03:41:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// id
	//
	// example:
	//
	// 393
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The API key status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAIDBClusterApiKeyResponseBodyApiKey) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterApiKeyResponseBodyApiKey) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) GetDescription() *string {
	return s.Description
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) GetId() *string {
	return s.Id
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) GetStatus() *string {
	return s.Status
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) SetApiKey(v string) *CreateAIDBClusterApiKeyResponseBodyApiKey {
	s.ApiKey = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) SetCreateTime(v string) *CreateAIDBClusterApiKeyResponseBodyApiKey {
	s.CreateTime = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) SetDescription(v string) *CreateAIDBClusterApiKeyResponseBodyApiKey {
	s.Description = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) SetId(v string) *CreateAIDBClusterApiKeyResponseBodyApiKey {
	s.Id = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) SetStatus(v string) *CreateAIDBClusterApiKeyResponseBodyApiKey {
	s.Status = &v
	return s
}

func (s *CreateAIDBClusterApiKeyResponseBodyApiKey) Validate() error {
	return dara.Validate(s)
}
