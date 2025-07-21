// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DedicatedIpPoolCreateResponseBody
	GetId() *string
	SetRequestId(v string) *DedicatedIpPoolCreateResponseBody
	GetRequestId() *string
}

type DedicatedIpPoolCreateResponseBody struct {
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

func (s DedicatedIpPoolCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolCreateResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolCreateResponseBody) GetId() *string {
	return s.Id
}

func (s *DedicatedIpPoolCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpPoolCreateResponseBody) SetId(v string) *DedicatedIpPoolCreateResponseBody {
	s.Id = &v
	return s
}

func (s *DedicatedIpPoolCreateResponseBody) SetRequestId(v string) *DedicatedIpPoolCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpPoolCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
