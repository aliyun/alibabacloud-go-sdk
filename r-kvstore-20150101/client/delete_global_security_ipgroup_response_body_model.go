// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGlobalSecurityIPGroupResponseBody
	GetRequestId() *string
}

type DeleteGlobalSecurityIPGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD425AD3-CC7B-4EE2-A5CB-2F61BA73****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) SetRequestId(v string) *DeleteGlobalSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
