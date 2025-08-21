// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionExternalApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeprovisionExternalApplicationResponseBody
	GetRequestId() *string
}

type DeprovisionExternalApplicationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1C2-3AB045D1AE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeprovisionExternalApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionExternalApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeprovisionExternalApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeprovisionExternalApplicationResponseBody) SetRequestId(v string) *DeprovisionExternalApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeprovisionExternalApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
