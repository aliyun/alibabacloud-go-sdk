// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryInstanceResponseBody
	GetData() *string
	SetRequestId(v string) *QueryInstanceResponseBody
	GetRequestId() *string
}

type QueryInstanceResponseBody struct {
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceResponseBody) SetData(v string) *QueryInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *QueryInstanceResponseBody) SetRequestId(v string) *QueryInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
