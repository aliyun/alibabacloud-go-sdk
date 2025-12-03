// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkRelateHBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SparkRelateHBaseResponseBody
	GetRequestId() *string
}

type SparkRelateHBaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SparkRelateHBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SparkRelateHBaseResponseBody) GoString() string {
	return s.String()
}

func (s *SparkRelateHBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SparkRelateHBaseResponseBody) SetRequestId(v string) *SparkRelateHBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SparkRelateHBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
