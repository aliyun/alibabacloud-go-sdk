// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSCIMServerCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSCIMServerCredentialResponseBody
	GetRequestId() *string
}

type DeleteSCIMServerCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8CE8B990-193D-50CE-A604-69F3E7DCE740
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSCIMServerCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSCIMServerCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSCIMServerCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSCIMServerCredentialResponseBody) SetRequestId(v string) *DeleteSCIMServerCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSCIMServerCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
