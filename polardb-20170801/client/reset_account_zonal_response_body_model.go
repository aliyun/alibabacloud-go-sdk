// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAccountZonalResponseBody
	GetRequestId() *string
}

type ResetAccountZonalResponseBody struct {
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountZonalResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountZonalResponseBody) SetRequestId(v string) *ResetAccountZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
