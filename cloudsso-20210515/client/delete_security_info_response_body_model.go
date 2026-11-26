// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityInfoResponseBody
	GetRequestId() *string
}

type DeleteSecurityInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B9982ED-FD0D-5622-8EA0-7B768685DCE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityInfoResponseBody) SetRequestId(v string) *DeleteSecurityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
