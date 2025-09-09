// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcPeerConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcPeerConnectionResponseBody
	GetRequestId() *string
}

type ModifyVpcPeerConnectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 880C99E1-449B-524A-B81F-1EC53D2A7EAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcPeerConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcPeerConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcPeerConnectionResponseBody) SetRequestId(v string) *ModifyVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcPeerConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
