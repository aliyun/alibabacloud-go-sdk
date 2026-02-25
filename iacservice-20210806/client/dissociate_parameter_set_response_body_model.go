// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateParameterSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateParameterSetResponseBody
	GetRequestId() *string
}

type DissociateParameterSetResponseBody struct {
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DissociateParameterSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateParameterSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateParameterSetResponseBody) SetRequestId(v string) *DissociateParameterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateParameterSetResponseBody) Validate() error {
	return dara.Validate(s)
}
