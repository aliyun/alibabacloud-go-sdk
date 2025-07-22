// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisFactorDistributionStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeFaultDiagnosisFactorDistributionStatRequest
	GetAppId() *string
	SetEndTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest
	GetEndTs() *int64
	SetStartTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest
	GetStartTs() *int64
}

type DescribeFaultDiagnosisFactorDistributionStatRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596****
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetAppId(v string) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetEndTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetStartTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) Validate() error {
	return dara.Validate(s)
}
