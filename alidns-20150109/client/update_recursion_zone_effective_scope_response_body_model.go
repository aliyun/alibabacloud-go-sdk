// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneEffectiveScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionZoneEffectiveScopeResponseBody
	GetRequestId() *string
}

type UpdateRecursionZoneEffectiveScopeResponseBody struct {
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionZoneEffectiveScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneEffectiveScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneEffectiveScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionZoneEffectiveScopeResponseBody) SetRequestId(v string) *UpdateRecursionZoneEffectiveScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
