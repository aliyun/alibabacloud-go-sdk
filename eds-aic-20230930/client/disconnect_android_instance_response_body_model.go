// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisconnectAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisconnectAndroidInstanceResponseBody
	GetRequestId() *string
}

type DisconnectAndroidInstanceResponseBody struct {
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisconnectAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisconnectAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DisconnectAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisconnectAndroidInstanceResponseBody) SetRequestId(v string) *DisconnectAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisconnectAndroidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
