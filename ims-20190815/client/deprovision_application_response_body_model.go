// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeprovisionApplicationResponseBody
	GetRequestId() *string
}

type DeprovisionApplicationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeprovisionApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeprovisionApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeprovisionApplicationResponseBody) SetRequestId(v string) *DeprovisionApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeprovisionApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
