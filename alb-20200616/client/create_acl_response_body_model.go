// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateAclResponseBody
	GetAclId() *string
	SetJobId(v string) *CreateAclResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateAclResponseBody
	GetRequestId() *string
}

type CreateAclResponseBody struct {
	// The ID of the ACL.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *CreateAclResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAclResponseBody) SetAclId(v string) *CreateAclResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAclResponseBody) SetJobId(v string) *CreateAclResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclResponseBody) Validate() error {
	return dara.Validate(s)
}
