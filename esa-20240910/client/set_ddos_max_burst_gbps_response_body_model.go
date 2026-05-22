// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDdosMaxBurstGbpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDdosMaxBurstGbpsResponseBody
	GetRequestId() *string
}

type SetDdosMaxBurstGbpsResponseBody struct {
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDdosMaxBurstGbpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDdosMaxBurstGbpsResponseBody) GoString() string {
	return s.String()
}

func (s *SetDdosMaxBurstGbpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDdosMaxBurstGbpsResponseBody) SetRequestId(v string) *SetDdosMaxBurstGbpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDdosMaxBurstGbpsResponseBody) Validate() error {
	return dara.Validate(s)
}
