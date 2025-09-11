// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckOrganizationMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckOrganizationMemberResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckOrganizationMemberResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CheckOrganizationMemberResponseBody
	GetSuccess() *bool
}

type CheckOrganizationMemberResponseBody struct {
	// example:
	//
	// D787E1A**********DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckOrganizationMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CheckOrganizationMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckOrganizationMemberResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckOrganizationMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckOrganizationMemberResponseBody) SetRequestId(v string) *CheckOrganizationMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckOrganizationMemberResponseBody) SetResult(v bool) *CheckOrganizationMemberResponseBody {
	s.Result = &v
	return s
}

func (s *CheckOrganizationMemberResponseBody) SetSuccess(v bool) *CheckOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

func (s *CheckOrganizationMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
