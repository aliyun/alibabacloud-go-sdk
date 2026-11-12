// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageShareStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyImageShareStatusResponseBody
	GetRequestId() *string
}

type ModifyImageShareStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageShareStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageShareStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageShareStatusResponseBody) SetRequestId(v string) *ModifyImageShareStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageShareStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
