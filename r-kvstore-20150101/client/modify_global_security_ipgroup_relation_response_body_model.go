// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupRelationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1E83311F-0EE4-4922-A3BF-730B312B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
