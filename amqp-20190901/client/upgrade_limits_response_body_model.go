// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLimitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpgradeLimitsResponseBody
	GetCode() *int32
	SetData(v string) *UpgradeLimitsResponseBody
	GetData() *string
	SetMessage(v string) *UpgradeLimitsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeLimitsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeLimitsResponseBody
	GetSuccess() *bool
}

type UpgradeLimitsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeLimitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeLimitsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradeLimitsResponseBody) GetData() *string {
	return s.Data
}

func (s *UpgradeLimitsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeLimitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeLimitsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeLimitsResponseBody) SetCode(v int32) *UpgradeLimitsResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeLimitsResponseBody) SetData(v string) *UpgradeLimitsResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeLimitsResponseBody) SetMessage(v string) *UpgradeLimitsResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeLimitsResponseBody) SetRequestId(v string) *UpgradeLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeLimitsResponseBody) SetSuccess(v bool) *UpgradeLimitsResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeLimitsResponseBody) Validate() error {
	return dara.Validate(s)
}
