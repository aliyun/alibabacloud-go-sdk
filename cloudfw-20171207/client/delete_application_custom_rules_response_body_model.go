// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationCustomRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationCustomRulesResponseBody
	GetRequestId() *string
}

type DeleteApplicationCustomRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationCustomRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationCustomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationCustomRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationCustomRulesResponseBody) SetRequestId(v string) *DeleteApplicationCustomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationCustomRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
