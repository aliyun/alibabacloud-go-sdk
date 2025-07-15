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
	// The request ID.
	//
	// example:
	//
	// 45B7AA4A-4658-5FFC-90DD-9B8729F301BB
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
