// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKillInstanceSessionTaskWithMaintainUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody
	GetCode() *int64
	SetData(v string) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody
	GetData() *string
	SetMessage(v string) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody
	GetSuccess() *bool
}

type CreateKillInstanceSessionTaskWithMaintainUserResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKillInstanceSessionTaskWithMaintainUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKillInstanceSessionTaskWithMaintainUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) SetCode(v int64) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) SetData(v string) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) SetMessage(v string) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) SetRequestId(v string) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) SetSuccess(v bool) *CreateKillInstanceSessionTaskWithMaintainUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) Validate() error {
	return dara.Validate(s)
}
