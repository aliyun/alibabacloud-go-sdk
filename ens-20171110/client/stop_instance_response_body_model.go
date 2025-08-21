// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *StopInstanceResponseBody
	GetCode() *int32
	SetRequestId(v string) *StopInstanceResponseBody
	GetRequestId() *string
}

type StopInstanceResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *StopInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstanceResponseBody) SetCode(v int32) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
