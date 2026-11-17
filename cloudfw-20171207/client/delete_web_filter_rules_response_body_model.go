// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebFilterRulesResponseBody
	GetRequestId() *string
}

type DeleteWebFilterRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebFilterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebFilterRulesResponseBody) SetRequestId(v string) *DeleteWebFilterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebFilterRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
