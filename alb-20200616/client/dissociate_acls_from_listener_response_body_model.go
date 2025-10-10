// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAclsFromListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DissociateAclsFromListenerResponseBody
	GetJobId() *string
	SetRequestId(v string) *DissociateAclsFromListenerResponseBody
	GetRequestId() *string
}

type DissociateAclsFromListenerResponseBody struct {
	// The asynchronous task ID.
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

func (s DissociateAclsFromListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateAclsFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DissociateAclsFromListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateAclsFromListenerResponseBody) SetJobId(v string) *DissociateAclsFromListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetRequestId(v string) *DissociateAclsFromListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
