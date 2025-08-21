// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShrinkNodeResponseBody
	GetRequestId() *string
	SetResult(v bool) *ShrinkNodeResponseBody
	GetResult() *bool
}

type ShrinkNodeResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ShrinkNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShrinkNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ShrinkNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShrinkNodeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ShrinkNodeResponseBody) SetRequestId(v string) *ShrinkNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShrinkNodeResponseBody) SetResult(v bool) *ShrinkNodeResponseBody {
	s.Result = &v
	return s
}

func (s *ShrinkNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
