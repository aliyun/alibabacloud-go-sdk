// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiServiceProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAiServiceProtectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAiServiceProtectionResponseBody
	GetSuccess() *bool
}

type ModifyAiServiceProtectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAiServiceProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiServiceProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAiServiceProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAiServiceProtectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAiServiceProtectionResponseBody) SetRequestId(v string) *ModifyAiServiceProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAiServiceProtectionResponseBody) SetSuccess(v bool) *ModifyAiServiceProtectionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAiServiceProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
