// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBagRemainingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *QueryBagRemainingResponseBody
	GetData() *int64
	SetRequestId(v string) *QueryBagRemainingResponseBody
	GetRequestId() *string
}

type QueryBagRemainingResponseBody struct {
	// example:
	//
	// True
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 6A603AA0-73BA-52B3-AC7D-0F846ECF7A9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryBagRemainingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBagRemainingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBagRemainingResponseBody) GetData() *int64 {
	return s.Data
}

func (s *QueryBagRemainingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBagRemainingResponseBody) SetData(v int64) *QueryBagRemainingResponseBody {
	s.Data = &v
	return s
}

func (s *QueryBagRemainingResponseBody) SetRequestId(v string) *QueryBagRemainingResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBagRemainingResponseBody) Validate() error {
	return dara.Validate(s)
}
