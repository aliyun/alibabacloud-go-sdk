// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccountZonalResponseBody
	GetRequestId() *string
}

type DeleteAccountZonalResponseBody struct {
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccountZonalResponseBody) SetRequestId(v string) *DeleteAccountZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
