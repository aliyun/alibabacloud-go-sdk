// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlId(v string) *CreateIpControlResponseBody
	GetIpControlId() *string
	SetRequestId(v string) *CreateIpControlResponseBody
	GetRequestId() *string
}

type CreateIpControlResponseBody struct {
	// The ID of the ACL.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CE5722A6-AE78-4741-A9B0-6C817D360510
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpControlResponseBody) GetIpControlId() *string {
	return s.IpControlId
}

func (s *CreateIpControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpControlResponseBody) SetIpControlId(v string) *CreateIpControlResponseBody {
	s.IpControlId = &v
	return s
}

func (s *CreateIpControlResponseBody) SetRequestId(v string) *CreateIpControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpControlResponseBody) Validate() error {
	return dara.Validate(s)
}
