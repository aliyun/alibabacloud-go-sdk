// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPenetrationTestDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *QueryPenetrationTestDetailRequest
	GetId() *int64
}

type QueryPenetrationTestDetailRequest struct {
	// The unique ID of the application record.
	//
	// > This field is of the Long type. Precision loss may occur when you serialize or deserialize the value. Make sure that the value does not exceed 9007199254740991.
	//
	// example:
	//
	// 2398
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryPenetrationTestDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryPenetrationTestDetailRequest) SetId(v int64) *QueryPenetrationTestDetailRequest {
	s.Id = &v
	return s
}

func (s *QueryPenetrationTestDetailRequest) Validate() error {
	return dara.Validate(s)
}
