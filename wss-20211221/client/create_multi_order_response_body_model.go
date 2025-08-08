// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderIds(v []*int64) *CreateMultiOrderResponseBody
	GetOrderIds() []*int64
	SetRequestId(v string) *CreateMultiOrderResponseBody
	GetRequestId() *string
}

type CreateMultiOrderResponseBody struct {
	OrderIds []*int64 `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderResponseBody) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *CreateMultiOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiOrderResponseBody) SetOrderIds(v []*int64) *CreateMultiOrderResponseBody {
	s.OrderIds = v
	return s
}

func (s *CreateMultiOrderResponseBody) SetRequestId(v string) *CreateMultiOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
