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
	// The ID of the request. If you encounter an issue, provide this request ID to our support staff for troubleshooting.
	//
	// example:
	//
	// 2F42BB4E-461F-5B55-A37C-53B1141C****
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
