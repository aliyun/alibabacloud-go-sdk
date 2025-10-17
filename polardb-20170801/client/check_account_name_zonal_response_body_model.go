// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckAccountNameZonalResponseBody
	GetRequestId() *string
}

type CheckAccountNameZonalResponseBody struct {
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountNameZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameZonalResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountNameZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountNameZonalResponseBody) SetRequestId(v string) *CheckAccountNameZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountNameZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
