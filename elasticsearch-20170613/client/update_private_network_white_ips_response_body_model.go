// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateNetworkWhiteIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePrivateNetworkWhiteIpsResponseBody
	GetRequestId() *string
	SetResult(v *UpdatePrivateNetworkWhiteIpsResponseBodyResult) *UpdatePrivateNetworkWhiteIpsResponseBody
	GetResult() *UpdatePrivateNetworkWhiteIpsResponseBodyResult
}

type UpdatePrivateNetworkWhiteIpsResponseBody struct {
	// example:
	//
	// 6DEBE5EE-0368-4757-8F82-EF9C3972****
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdatePrivateNetworkWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdatePrivateNetworkWhiteIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) GetResult() *UpdatePrivateNetworkWhiteIpsResponseBodyResult {
	return s.Result
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) SetRequestId(v string) *UpdatePrivateNetworkWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) SetResult(v *UpdatePrivateNetworkWhiteIpsResponseBodyResult) *UpdatePrivateNetworkWhiteIpsResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePrivateNetworkWhiteIpsResponseBodyResult struct {
	PrivateNetworkIpWhiteList []*string `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetworkWhiteIpsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBodyResult) GetPrivateNetworkIpWhiteList() []*string {
	return s.PrivateNetworkIpWhiteList
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBodyResult) SetPrivateNetworkIpWhiteList(v []*string) *UpdatePrivateNetworkWhiteIpsResponseBodyResult {
	s.PrivateNetworkIpWhiteList = v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
