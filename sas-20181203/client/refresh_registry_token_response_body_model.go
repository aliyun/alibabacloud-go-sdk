// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRegistryTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RefreshRegistryTokenResponseBody
	GetData() *string
	SetRequestId(v string) *RefreshRegistryTokenResponseBody
	GetRequestId() *string
}

type RefreshRegistryTokenResponseBody struct {
	// The returned token.
	//
	// example:
	//
	// 77ba3bf5-af95-4b77-aa94-***********
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshRegistryTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshRegistryTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshRegistryTokenResponseBody) GetData() *string {
	return s.Data
}

func (s *RefreshRegistryTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshRegistryTokenResponseBody) SetData(v string) *RefreshRegistryTokenResponseBody {
	s.Data = &v
	return s
}

func (s *RefreshRegistryTokenResponseBody) SetRequestId(v string) *RefreshRegistryTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshRegistryTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
