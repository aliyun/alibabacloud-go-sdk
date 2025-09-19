// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartIdcProbeScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartIdcProbeScanResponseBody
	GetRequestId() *string
}

type StartIdcProbeScanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53CXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartIdcProbeScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartIdcProbeScanResponseBody) GoString() string {
	return s.String()
}

func (s *StartIdcProbeScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartIdcProbeScanResponseBody) SetRequestId(v string) *StartIdcProbeScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartIdcProbeScanResponseBody) Validate() error {
	return dara.Validate(s)
}
