// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNetworkRulesResponseBody
	GetRequestId() *string
}

type CreateNetworkRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ADCA45A5-D15C-4B7D-9F81-138B0B36D0BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkRulesResponseBody) SetRequestId(v string) *CreateNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
