// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnsRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEnsRouteEntryResponseBody
	GetRequestId() *string
}

type ModifyEnsRouteEntryResponseBody struct {
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEnsRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnsRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEnsRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEnsRouteEntryResponseBody) SetRequestId(v string) *ModifyEnsRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEnsRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
