// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceCredentialResponseBody
	GetRequestId() *string
}

type DeleteServiceCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 492E130C-76D3-55D5-BE5C-C023E431369A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceCredentialResponseBody) SetRequestId(v string) *DeleteServiceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
