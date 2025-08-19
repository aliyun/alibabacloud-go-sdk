// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDictResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDictResponseBody
	GetResult() *bool
}

type UpdateDictResponseBody struct {
	// example:
	//
	// 12797BCC-E0B5-5A47-B4B9-A14DDF0B0200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDictResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDictResponseBody) SetRequestId(v string) *UpdateDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDictResponseBody) SetResult(v bool) *UpdateDictResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDictResponseBody) Validate() error {
	return dara.Validate(s)
}
