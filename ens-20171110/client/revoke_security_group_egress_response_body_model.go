// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupEgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeSecurityGroupEgressResponseBody
	GetRequestId() *string
}

type RevokeSecurityGroupEgressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeSecurityGroupEgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupEgressResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeSecurityGroupEgressResponseBody) SetRequestId(v string) *RevokeSecurityGroupEgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeSecurityGroupEgressResponseBody) Validate() error {
	return dara.Validate(s)
}
