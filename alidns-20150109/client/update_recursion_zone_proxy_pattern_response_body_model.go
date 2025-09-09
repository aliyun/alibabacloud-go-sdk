// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneProxyPatternResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionZoneProxyPatternResponseBody
	GetRequestId() *string
}

type UpdateRecursionZoneProxyPatternResponseBody struct {
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionZoneProxyPatternResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneProxyPatternResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneProxyPatternResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionZoneProxyPatternResponseBody) SetRequestId(v string) *UpdateRecursionZoneProxyPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionZoneProxyPatternResponseBody) Validate() error {
	return dara.Validate(s)
}
