// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAuthModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVpcAuthModeResponseBody
	GetRequestId() *string
}

type ModifyInstanceVpcAuthModeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABAF95F6-35C1-4177-AF3A-70969EBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVpcAuthModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVpcAuthModeResponseBody) SetRequestId(v string) *ModifyInstanceVpcAuthModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeResponseBody) Validate() error {
	return dara.Validate(s)
}
