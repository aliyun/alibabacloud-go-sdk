// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWebFilterRulesResponseBody
	GetRequestId() *string
}

type UpdateWebFilterRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebFilterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWebFilterRulesResponseBody) SetRequestId(v string) *UpdateWebFilterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWebFilterRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
