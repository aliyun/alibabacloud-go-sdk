// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecalculateCarbonEmissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RecalculateCarbonEmissionResponseBody
	GetData() *bool
	SetRequestId(v string) *RecalculateCarbonEmissionResponseBody
	GetRequestId() *string
}

type RecalculateCarbonEmissionResponseBody struct {
	// The returned data. A value of true indicates that the request is successful. A value of false indicates that the request fails.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RecalculateCarbonEmissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecalculateCarbonEmissionResponseBody) GoString() string {
	return s.String()
}

func (s *RecalculateCarbonEmissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *RecalculateCarbonEmissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecalculateCarbonEmissionResponseBody) SetData(v bool) *RecalculateCarbonEmissionResponseBody {
	s.Data = &v
	return s
}

func (s *RecalculateCarbonEmissionResponseBody) SetRequestId(v string) *RecalculateCarbonEmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecalculateCarbonEmissionResponseBody) Validate() error {
	return dara.Validate(s)
}
