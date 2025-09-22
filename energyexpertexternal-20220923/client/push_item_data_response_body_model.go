// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushItemDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *PushItemDataResponseBody
	GetData() *bool
	SetRequestId(v string) *PushItemDataResponseBody
	GetRequestId() *string
}

type PushItemDataResponseBody struct {
	// Whether the data is pushed successfully.
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

func (s PushItemDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushItemDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushItemDataResponseBody) GetData() *bool {
	return s.Data
}

func (s *PushItemDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushItemDataResponseBody) SetData(v bool) *PushItemDataResponseBody {
	s.Data = &v
	return s
}

func (s *PushItemDataResponseBody) SetRequestId(v string) *PushItemDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushItemDataResponseBody) Validate() error {
	return dara.Validate(s)
}
