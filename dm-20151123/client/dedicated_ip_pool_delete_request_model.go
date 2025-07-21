// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DedicatedIpPoolDeleteRequest
	GetId() *string
}

type DedicatedIpPoolDeleteRequest struct {
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DedicatedIpPoolDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolDeleteRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolDeleteRequest) GetId() *string {
	return s.Id
}

func (s *DedicatedIpPoolDeleteRequest) SetId(v string) *DedicatedIpPoolDeleteRequest {
	s.Id = &v
	return s
}

func (s *DedicatedIpPoolDeleteRequest) Validate() error {
	return dara.Validate(s)
}
