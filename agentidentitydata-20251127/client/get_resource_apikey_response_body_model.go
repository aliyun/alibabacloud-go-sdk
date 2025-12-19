// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceAPIKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *GetResourceAPIKeyResponseBody
	GetAPIKey() *string
	SetRequestId(v string) *GetResourceAPIKeyResponseBody
	GetRequestId() *string
}

type GetResourceAPIKeyResponseBody struct {
	// example:
	//
	// sk-ds*****
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// example:
	//
	// 6E444C1B-1106-56A8-81E0-E3B049BF44AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceAPIKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceAPIKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceAPIKeyResponseBody) GetAPIKey() *string {
	return s.APIKey
}

func (s *GetResourceAPIKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceAPIKeyResponseBody) SetAPIKey(v string) *GetResourceAPIKeyResponseBody {
	s.APIKey = &v
	return s
}

func (s *GetResourceAPIKeyResponseBody) SetRequestId(v string) *GetResourceAPIKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceAPIKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
