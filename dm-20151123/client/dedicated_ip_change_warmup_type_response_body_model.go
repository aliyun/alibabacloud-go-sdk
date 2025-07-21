// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpChangeWarmupTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DedicatedIpChangeWarmupTypeResponseBody
	GetRequestId() *string
}

type DedicatedIpChangeWarmupTypeResponseBody struct {
	// Request ID
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DedicatedIpChangeWarmupTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpChangeWarmupTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpChangeWarmupTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpChangeWarmupTypeResponseBody) SetRequestId(v string) *DedicatedIpChangeWarmupTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpChangeWarmupTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
