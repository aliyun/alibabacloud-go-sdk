// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleConditionControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnsSaleConditionControlResponseBody
	GetRequestId() *string
}

type DeleteEnsSaleConditionControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnsSaleConditionControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleConditionControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleConditionControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnsSaleConditionControlResponseBody) SetRequestId(v string) *DeleteEnsSaleConditionControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnsSaleConditionControlResponseBody) Validate() error {
	return dara.Validate(s)
}
