// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountFactoryBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccountFactoryBaselineResponseBody
	GetRequestId() *string
}

type DeleteAccountFactoryBaselineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0F45D888-8C4D-55E5-ACA2-D1515159181D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountFactoryBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountFactoryBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccountFactoryBaselineResponseBody) SetRequestId(v string) *DeleteAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountFactoryBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
