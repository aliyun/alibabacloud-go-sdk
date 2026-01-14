// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAclsFromListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *DissociateAclsFromListenerResponseBody
	GetAclIds() []*string
	SetListenerId(v string) *DissociateAclsFromListenerResponseBody
	GetListenerId() *string
	SetRequestId(v string) *DissociateAclsFromListenerResponseBody
	GetRequestId() *string
}

type DissociateAclsFromListenerResponseBody struct {
	// The IDs of the ACL.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the request.
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

func (s *DissociateAclsFromListenerResponseBody) GetAclIds() []*string {
	return s.AclIds
}

func (s *DissociateAclsFromListenerResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DissociateAclsFromListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateAclsFromListenerResponseBody) SetAclIds(v []*string) *DissociateAclsFromListenerResponseBody {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetListenerId(v string) *DissociateAclsFromListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetRequestId(v string) *DissociateAclsFromListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
