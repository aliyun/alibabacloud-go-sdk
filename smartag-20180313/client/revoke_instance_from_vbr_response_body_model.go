// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromVbrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeInstanceFromVbrResponseBody
	GetRequestId() *string
}

type RevokeInstanceFromVbrResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D085AE49-51DC-4E8A-9F06-2D99C4E374F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeInstanceFromVbrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromVbrResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromVbrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeInstanceFromVbrResponseBody) SetRequestId(v string) *RevokeInstanceFromVbrResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromVbrResponseBody) Validate() error {
	return dara.Validate(s)
}
