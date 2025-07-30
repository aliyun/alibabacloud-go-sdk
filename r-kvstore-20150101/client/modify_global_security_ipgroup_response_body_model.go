// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 686BB8A6-BBA5-47E5-8A75-D2ADE433****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
