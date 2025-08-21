// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RebootInstanceResponseBody
	GetCode() *int32
	SetRequestId(v string) *RebootInstanceResponseBody
	GetRequestId() *string
}

type RebootInstanceResponseBody struct {
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

func (s RebootInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RebootInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootInstanceResponseBody) SetCode(v int32) *RebootInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RebootInstanceResponseBody) SetRequestId(v string) *RebootInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
