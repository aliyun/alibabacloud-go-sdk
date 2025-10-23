// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetRelationFromAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigSetRelationFromAddressResponseBody
	GetRequestId() *string
	SetResult(v bool) *ConfigSetRelationFromAddressResponseBody
	GetResult() *bool
}

type ConfigSetRelationFromAddressResponseBody struct {
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ConfigSetRelationFromAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetRelationFromAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetRelationFromAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetRelationFromAddressResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ConfigSetRelationFromAddressResponseBody) SetRequestId(v string) *ConfigSetRelationFromAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetRelationFromAddressResponseBody) SetResult(v bool) *ConfigSetRelationFromAddressResponseBody {
	s.Result = &v
	return s
}

func (s *ConfigSetRelationFromAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
