// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEgressProxyRulesResponseBody
	GetRequestId() *string
}

type DeleteEgressProxyRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C79118A9-CA6D-5359-B53E-111BE8A03DCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEgressProxyRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEgressProxyRulesResponseBody) SetRequestId(v string) *DeleteEgressProxyRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEgressProxyRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
