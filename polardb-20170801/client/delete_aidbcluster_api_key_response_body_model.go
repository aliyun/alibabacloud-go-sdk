// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DeleteAIDBClusterApiKeyResponseBody
	GetApiKey() *string
	SetRequestId(v string) *DeleteAIDBClusterApiKeyResponseBody
	GetRequestId() *string
}

type DeleteAIDBClusterApiKeyResponseBody struct {
	// The API key of the model service.
	//
	// example:
	//
	// xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 068F730C-9130-596E-B696-5B4388C840DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIDBClusterApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeleteAIDBClusterApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIDBClusterApiKeyResponseBody) SetApiKey(v string) *DeleteAIDBClusterApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *DeleteAIDBClusterApiKeyResponseBody) SetRequestId(v string) *DeleteAIDBClusterApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIDBClusterApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
