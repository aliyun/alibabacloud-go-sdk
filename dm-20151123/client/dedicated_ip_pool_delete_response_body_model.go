// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DedicatedIpPoolDeleteResponseBody
	GetRequestId() *string
}

type DedicatedIpPoolDeleteResponseBody struct {
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DedicatedIpPoolDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpPoolDeleteResponseBody) SetRequestId(v string) *DedicatedIpPoolDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpPoolDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
