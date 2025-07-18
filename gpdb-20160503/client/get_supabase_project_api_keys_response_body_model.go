// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*GetSupabaseProjectApiKeysResponseBodyApiKeys) *GetSupabaseProjectApiKeysResponseBody
	GetApiKeys() []*GetSupabaseProjectApiKeysResponseBodyApiKeys
	SetRequestId(v string) *GetSupabaseProjectApiKeysResponseBody
	GetRequestId() *string
}

type GetSupabaseProjectApiKeysResponseBody struct {
	ApiKeys []*GetSupabaseProjectApiKeysResponseBodyApiKeys `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSupabaseProjectApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectApiKeysResponseBody) GetApiKeys() []*GetSupabaseProjectApiKeysResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *GetSupabaseProjectApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupabaseProjectApiKeysResponseBody) SetApiKeys(v []*GetSupabaseProjectApiKeysResponseBodyApiKeys) *GetSupabaseProjectApiKeysResponseBody {
	s.ApiKeys = v
	return s
}

func (s *GetSupabaseProjectApiKeysResponseBody) SetRequestId(v string) *GetSupabaseProjectApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupabaseProjectApiKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSupabaseProjectApiKeysResponseBodyApiKeys struct {
	// example:
	//
	// Tmz2Z59caMDeq/Xi9vuc****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// anon key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupabaseProjectApiKeysResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectApiKeysResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectApiKeysResponseBodyApiKeys) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetSupabaseProjectApiKeysResponseBodyApiKeys) GetName() *string {
	return s.Name
}

func (s *GetSupabaseProjectApiKeysResponseBodyApiKeys) SetApiKey(v string) *GetSupabaseProjectApiKeysResponseBodyApiKeys {
	s.ApiKey = &v
	return s
}

func (s *GetSupabaseProjectApiKeysResponseBodyApiKeys) SetName(v string) *GetSupabaseProjectApiKeysResponseBodyApiKeys {
	s.Name = &v
	return s
}

func (s *GetSupabaseProjectApiKeysResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}
