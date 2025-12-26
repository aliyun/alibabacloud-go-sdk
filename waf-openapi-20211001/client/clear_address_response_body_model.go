// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearAddressResponseBody
	GetRequestId() *string
}

type ClearAddressResponseBody struct {
	// example:
	//
	// 276D7566-31C9-4192-9DD1-51B10D*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearAddressResponseBody) SetRequestId(v string) *ClearAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
