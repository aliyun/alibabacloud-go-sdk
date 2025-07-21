// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DedicatedIpPoolUpdateResponseBody
	GetId() *string
	SetRequestId(v string) *DedicatedIpPoolUpdateResponseBody
	GetRequestId() *string
}

type DedicatedIpPoolUpdateResponseBody struct {
	// IP pool ID
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Request ID
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DedicatedIpPoolUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolUpdateResponseBody) GetId() *string {
	return s.Id
}

func (s *DedicatedIpPoolUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpPoolUpdateResponseBody) SetId(v string) *DedicatedIpPoolUpdateResponseBody {
	s.Id = &v
	return s
}

func (s *DedicatedIpPoolUpdateResponseBody) SetRequestId(v string) *DedicatedIpPoolUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpPoolUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
