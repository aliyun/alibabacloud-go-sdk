// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAclsWithListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *AssociateAclsWithListenerResponseBody
	GetAclIds() []*string
	SetListenerId(v string) *AssociateAclsWithListenerResponseBody
	GetListenerId() *string
	SetRequestId(v string) *AssociateAclsWithListenerResponseBody
	GetRequestId() *string
}

type AssociateAclsWithListenerResponseBody struct {
	// The ID of the ACL.
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
	// 64ADAB1E-0B7F-4FD8-A404-3BECC0E9CCFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAclsWithListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateAclsWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponseBody) GetAclIds() []*string {
	return s.AclIds
}

func (s *AssociateAclsWithListenerResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *AssociateAclsWithListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateAclsWithListenerResponseBody) SetAclIds(v []*string) *AssociateAclsWithListenerResponseBody {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetListenerId(v string) *AssociateAclsWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetRequestId(v string) *AssociateAclsWithListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
