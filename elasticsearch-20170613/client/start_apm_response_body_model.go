// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartApmResponseBody
	GetRequestId() *string
	SetResult(v bool) *StartApmResponseBody
	GetResult() *bool
}

type StartApmResponseBody struct {
	// example:
	//
	// 526F30AB-4A43-55BA-910F-ACD275FD5F14
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartApmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartApmResponseBody) GoString() string {
	return s.String()
}

func (s *StartApmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartApmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StartApmResponseBody) SetRequestId(v string) *StartApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApmResponseBody) SetResult(v bool) *StartApmResponseBody {
	s.Result = &v
	return s
}

func (s *StartApmResponseBody) Validate() error {
	return dara.Validate(s)
}
