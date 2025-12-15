// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeProxyResponseBody
	GetRequestId() *string
}

type UpgradeProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B830CDBB-91D9-5571-B6F4-35C76266****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeProxyResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeProxyResponseBody) SetRequestId(v string) *UpgradeProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
