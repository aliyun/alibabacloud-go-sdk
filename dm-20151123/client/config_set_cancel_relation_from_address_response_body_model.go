// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCancelRelationFromAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigSetCancelRelationFromAddressResponseBody
	GetRequestId() *string
	SetResult(v bool) *ConfigSetCancelRelationFromAddressResponseBody
	GetResult() *bool
}

type ConfigSetCancelRelationFromAddressResponseBody struct {
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ConfigSetCancelRelationFromAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCancelRelationFromAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetCancelRelationFromAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetCancelRelationFromAddressResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ConfigSetCancelRelationFromAddressResponseBody) SetRequestId(v string) *ConfigSetCancelRelationFromAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetCancelRelationFromAddressResponseBody) SetResult(v bool) *ConfigSetCancelRelationFromAddressResponseBody {
	s.Result = &v
	return s
}

func (s *ConfigSetCancelRelationFromAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
