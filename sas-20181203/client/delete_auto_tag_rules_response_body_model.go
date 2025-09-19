// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoTagRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAutoTagRulesResponseBody
	GetRequestId() *string
}

type DeleteAutoTagRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD380235-A0B8-540D-A0D5-D6288446****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoTagRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoTagRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoTagRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoTagRulesResponseBody) SetRequestId(v string) *DeleteAutoTagRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoTagRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
