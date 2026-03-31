// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePauseProtectionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPauseStatus(v int32) *DescribePauseProtectionStatusResponseBody
	GetPauseStatus() *int32
	SetRequestId(v string) *DescribePauseProtectionStatusResponseBody
	GetRequestId() *string
}

type DescribePauseProtectionStatusResponseBody struct {
	// Indicates whether WAF protection is paused.
	//
	// 	- **0**: indicates that WAF protection is not paused. This is the default value.
	//
	// 	- **1**: indicates that WAF protection is paused.
	//
	// example:
	//
	// 0
	PauseStatus *int32 `json:"PauseStatus,omitempty" xml:"PauseStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-****-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePauseProtectionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePauseProtectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePauseProtectionStatusResponseBody) GetPauseStatus() *int32 {
	return s.PauseStatus
}

func (s *DescribePauseProtectionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePauseProtectionStatusResponseBody) SetPauseStatus(v int32) *DescribePauseProtectionStatusResponseBody {
	s.PauseStatus = &v
	return s
}

func (s *DescribePauseProtectionStatusResponseBody) SetRequestId(v string) *DescribePauseProtectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePauseProtectionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
