// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoGroupingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAutoGroupingConfigResponseBody
	GetRequestId() *string
}

type UpdateAutoGroupingConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAutoGroupingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoGroupingConfigResponseBody) SetRequestId(v string) *UpdateAutoGroupingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoGroupingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
