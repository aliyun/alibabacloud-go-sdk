// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckServiceLinkedRoleResponseBody
	GetRequestId() *string
	SetServiceLinkedRoleExists(v bool) *CheckServiceLinkedRoleResponseBody
	GetServiceLinkedRoleExists() *bool
}

type CheckServiceLinkedRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BE0D8D2-6702-5639-A9C2-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether an SLR is created.
	//
	// example:
	//
	// True
	ServiceLinkedRoleExists *bool `json:"ServiceLinkedRoleExists,omitempty" xml:"ServiceLinkedRoleExists,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceLinkedRoleResponseBody) GetServiceLinkedRoleExists() *bool {
	return s.ServiceLinkedRoleExists
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetServiceLinkedRoleExists(v bool) *CheckServiceLinkedRoleResponseBody {
	s.ServiceLinkedRoleExists = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
