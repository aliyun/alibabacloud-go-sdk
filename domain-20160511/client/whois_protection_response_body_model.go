// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhoisProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WhoisProtectionResponseBody
	GetRequestId() *string
	SetResult(v int32) *WhoisProtectionResponseBody
	GetResult() *int32
}

type WhoisProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s WhoisProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WhoisProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *WhoisProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WhoisProtectionResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *WhoisProtectionResponseBody) SetRequestId(v string) *WhoisProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *WhoisProtectionResponseBody) SetResult(v int32) *WhoisProtectionResponseBody {
	s.Result = &v
	return s
}

func (s *WhoisProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
