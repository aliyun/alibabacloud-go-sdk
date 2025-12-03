// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelateDbForHBaseHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RelateDbForHBaseHaResponseBody
	GetRequestId() *string
}

type RelateDbForHBaseHaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RelateDbForHBaseHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RelateDbForHBaseHaResponseBody) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RelateDbForHBaseHaResponseBody) SetRequestId(v string) *RelateDbForHBaseHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RelateDbForHBaseHaResponseBody) Validate() error {
	return dara.Validate(s)
}
