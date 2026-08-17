// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlinkAiServiceFreeQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlinkAiFreeQuotaDTO(v *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) *GetFlinkAiServiceFreeQuotaResponseBody
	GetFlinkAiFreeQuotaDTO() *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO
	SetRequestId(v string) *GetFlinkAiServiceFreeQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFlinkAiServiceFreeQuotaResponseBody
	GetSuccess() *bool
}

type GetFlinkAiServiceFreeQuotaResponseBody struct {
	// The Flink AI free quota data transfer object.
	FlinkAiFreeQuotaDTO *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO `json:"FlinkAiFreeQuotaDTO,omitempty" xml:"FlinkAiFreeQuotaDTO,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C8DF2A5B-6FBA-5651-A3D4-960F3664****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFlinkAiServiceFreeQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceFreeQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) GetFlinkAiFreeQuotaDTO() *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO {
	return s.FlinkAiFreeQuotaDTO
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) SetFlinkAiFreeQuotaDTO(v *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) *GetFlinkAiServiceFreeQuotaResponseBody {
	s.FlinkAiFreeQuotaDTO = v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) SetRequestId(v string) *GetFlinkAiServiceFreeQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) SetSuccess(v bool) *GetFlinkAiServiceFreeQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBody) Validate() error {
	if s.FlinkAiFreeQuotaDTO != nil {
		if err := s.FlinkAiFreeQuotaDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO struct {
	// The total free quota.
	//
	// example:
	//
	// 100
	FreeQuota *float64 `json:"FreeQuota,omitempty" xml:"FreeQuota,omitempty"`
	// The list of used quota details for each usage type.
	UsedQuotaDetails []*GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails `json:"UsedQuotaDetails,omitempty" xml:"UsedQuotaDetails,omitempty" type:"Repeated"`
}

func (s GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) GetFreeQuota() *float64 {
	return s.FreeQuota
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) GetUsedQuotaDetails() []*GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails {
	return s.UsedQuotaDetails
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) SetFreeQuota(v float64) *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO {
	s.FreeQuota = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) SetUsedQuotaDetails(v []*GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO {
	s.UsedQuotaDetails = v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTO) Validate() error {
	if s.UsedQuotaDetails != nil {
		for _, item := range s.UsedQuotaDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails struct {
	// The used quota for this usage type.
	//
	// example:
	//
	// 35.5
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The usage type.
	//
	// example:
	//
	// AI_FUNCTION
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) GetAmount() *float64 {
	return s.Amount
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) GetUsageType() *string {
	return s.UsageType
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) SetAmount(v float64) *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails {
	s.Amount = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) SetUsageType(v string) *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails {
	s.UsageType = &v
	return s
}

func (s *GetFlinkAiServiceFreeQuotaResponseBodyFlinkAiFreeQuotaDTOUsedQuotaDetails) Validate() error {
	return dara.Validate(s)
}
