// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationCustomRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationCustomRulesResponseBody
	GetRequestId() *string
}

type UpdateApplicationCustomRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationCustomRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationCustomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationCustomRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationCustomRulesResponseBody) SetRequestId(v string) *UpdateApplicationCustomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationCustomRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
