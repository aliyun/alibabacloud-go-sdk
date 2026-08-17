// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlinkAiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlinkAiServiceDTO(v *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) *GetFlinkAiServiceResponseBody
	GetFlinkAiServiceDTO() *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO
	SetRequestId(v string) *GetFlinkAiServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFlinkAiServiceResponseBody
	GetSuccess() *bool
}

type GetFlinkAiServiceResponseBody struct {
	// The Flink AI service data transfer object.
	FlinkAiServiceDTO *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO `json:"FlinkAiServiceDTO,omitempty" xml:"FlinkAiServiceDTO,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B21DC47E-8928-199A-9F32-36D45E4693B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFlinkAiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceResponseBody) GetFlinkAiServiceDTO() *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO {
	return s.FlinkAiServiceDTO
}

func (s *GetFlinkAiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlinkAiServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFlinkAiServiceResponseBody) SetFlinkAiServiceDTO(v *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) *GetFlinkAiServiceResponseBody {
	s.FlinkAiServiceDTO = v
	return s
}

func (s *GetFlinkAiServiceResponseBody) SetRequestId(v string) *GetFlinkAiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlinkAiServiceResponseBody) SetSuccess(v bool) *GetFlinkAiServiceResponseBody {
	s.Success = &v
	return s
}

func (s *GetFlinkAiServiceResponseBody) Validate() error {
	if s.FlinkAiServiceDTO != nil {
		if err := s.FlinkAiServiceDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFlinkAiServiceResponseBodyFlinkAiServiceDTO struct {
	// Indicates whether deletion protection is enabled.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The status of the Flink AI instance. Valid values:
	//
	// - CLOSED: closed or not activated.
	//
	// - WAITING: waiting to be activated after payment.
	//
	// - OPENING: being activated.
	//
	// - RUNNING: activated.
	//
	// - CLOSING: being closed.
	//
	// - DISABLE: overdue payment.
	//
	// example:
	//
	// CLOSED
	FlinkAiInstanceStatus *string `json:"FlinkAiInstanceStatus,omitempty" xml:"FlinkAiInstanceStatus,omitempty"`
	// The AI service order instance ID.
	//
	// example:
	//
	// sc_flinkaifuncpost_public_cn-o3s4xabcdef
	MainInstanceId *string `json:"MainInstanceId,omitempty" xml:"MainInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the AI service was activated, in timestamp format (milliseconds).
	//
	// example:
	//
	// 1786934292
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
}

func (s GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) GetFlinkAiInstanceStatus() *string {
	return s.FlinkAiInstanceStatus
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) GetMainInstanceId() *string {
	return s.MainInstanceId
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) GetRegion() *string {
	return s.Region
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) GetResourceCreateTime() *int64 {
	return s.ResourceCreateTime
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) SetDeletionProtection(v bool) *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO {
	s.DeletionProtection = &v
	return s
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) SetFlinkAiInstanceStatus(v string) *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO {
	s.FlinkAiInstanceStatus = &v
	return s
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) SetMainInstanceId(v string) *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO {
	s.MainInstanceId = &v
	return s
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) SetRegion(v string) *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO {
	s.Region = &v
	return s
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) SetResourceCreateTime(v int64) *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetFlinkAiServiceResponseBodyFlinkAiServiceDTO) Validate() error {
	return dara.Validate(s)
}
