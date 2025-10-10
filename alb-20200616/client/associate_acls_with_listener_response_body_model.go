// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAclsWithListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *AssociateAclsWithListenerResponseBody
	GetJobId() *string
	SetRequestId(v string) *AssociateAclsWithListenerResponseBody
	GetRequestId() *string
}

type AssociateAclsWithListenerResponseBody struct {
	// The synchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAclsWithListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateAclsWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *AssociateAclsWithListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateAclsWithListenerResponseBody) SetJobId(v string) *AssociateAclsWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetRequestId(v string) *AssociateAclsWithListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
