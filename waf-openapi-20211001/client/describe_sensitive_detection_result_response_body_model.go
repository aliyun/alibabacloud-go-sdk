// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveDetectionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSensitiveDetectionResultResponseBodyData) *DescribeSensitiveDetectionResultResponseBody
	GetData() *DescribeSensitiveDetectionResultResponseBodyData
	SetRequestId(v string) *DescribeSensitiveDetectionResultResponseBody
	GetRequestId() *string
}

type DescribeSensitiveDetectionResultResponseBody struct {
	// The compliance check results.
	Data *DescribeSensitiveDetectionResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSensitiveDetectionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultResponseBody) GetData() *DescribeSensitiveDetectionResultResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveDetectionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveDetectionResultResponseBody) SetData(v *DescribeSensitiveDetectionResultResponseBodyData) *DescribeSensitiveDetectionResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBody) SetRequestId(v string) *DescribeSensitiveDetectionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSensitiveDetectionResultResponseBodyData struct {
	// The compliance checks.
	Result []*DescribeSensitiveDetectionResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeSensitiveDetectionResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultResponseBodyData) GetResult() []*DescribeSensitiveDetectionResultResponseBodyDataResult {
	return s.Result
}

func (s *DescribeSensitiveDetectionResultResponseBodyData) SetResult(v []*DescribeSensitiveDetectionResultResponseBodyDataResult) *DescribeSensitiveDetectionResultResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSensitiveDetectionResultResponseBodyDataResult struct {
	// The compliance check results. Valid values:
	//
	// 	- **report**: Risks exist in cross-border data transfer.
	//
	// 	- **none**: No risks exist in cross-border data transfer.
	//
	// example:
	//
	// report
	DetectionResult *string `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty"`
	// The sensitive information check results by sensitive data type.
	List []*DescribeSensitiveDetectionResultResponseBodyDataResultList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The maximum values in the statistics of sensitive data types.
	Max *DescribeSensitiveDetectionResultResponseBodyDataResultMax `json:"Max,omitempty" xml:"Max,omitempty" type:"Struct"`
}

func (s DescribeSensitiveDetectionResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) GetDetectionResult() *string {
	return s.DetectionResult
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) GetList() []*DescribeSensitiveDetectionResultResponseBodyDataResultList {
	return s.List
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) GetMax() *DescribeSensitiveDetectionResultResponseBodyDataResultMax {
	return s.Max
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) SetDetectionResult(v string) *DescribeSensitiveDetectionResultResponseBodyDataResult {
	s.DetectionResult = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) SetList(v []*DescribeSensitiveDetectionResultResponseBodyDataResultList) *DescribeSensitiveDetectionResultResponseBodyDataResult {
	s.List = v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) SetMax(v *DescribeSensitiveDetectionResultResponseBodyDataResultMax) *DescribeSensitiveDetectionResultResponseBodyDataResult {
	s.Max = v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResult) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Max != nil {
		if err := s.Max.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSensitiveDetectionResultResponseBodyDataResultList struct {
	// The number of personal information records.
	//
	// example:
	//
	// 11
	InfoCount *int64 `json:"InfoCount,omitempty" xml:"InfoCount,omitempty"`
	// The number of sensitive personal information records that are involved in cross-border data transfer.
	//
	// example:
	//
	// 6
	OutboundCount *int64 `json:"OutboundCount,omitempty" xml:"OutboundCount,omitempty"`
	// The sensitive data type.
	//
	// example:
	//
	// 1002
	SensitiveCode *int64 `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
}

func (s DescribeSensitiveDetectionResultResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) GetInfoCount() *int64 {
	return s.InfoCount
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) GetOutboundCount() *int64 {
	return s.OutboundCount
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) GetSensitiveCode() *int64 {
	return s.SensitiveCode
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) SetInfoCount(v int64) *DescribeSensitiveDetectionResultResponseBodyDataResultList {
	s.InfoCount = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) SetOutboundCount(v int64) *DescribeSensitiveDetectionResultResponseBodyDataResultList {
	s.OutboundCount = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) SetSensitiveCode(v int64) *DescribeSensitiveDetectionResultResponseBodyDataResultList {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultList) Validate() error {
	return dara.Validate(s)
}

type DescribeSensitiveDetectionResultResponseBodyDataResultMax struct {
	// The number of sensitive personal information records that are of the most frequent sensitive data type.
	//
	// example:
	//
	// 187
	InfoCount *int64 `json:"InfoCount,omitempty" xml:"InfoCount,omitempty"`
	// The number of sensitive personal information records that are of the most frequent sensitive data type and are involved in cross-border data transfer.
	//
	// example:
	//
	// 54
	OutboundCount *int64 `json:"OutboundCount,omitempty" xml:"OutboundCount,omitempty"`
	// The most frequent sensitive data type.
	//
	// example:
	//
	// 1003
	SensitiveCode *int64 `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
}

func (s DescribeSensitiveDetectionResultResponseBodyDataResultMax) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultResponseBodyDataResultMax) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) GetInfoCount() *int64 {
	return s.InfoCount
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) GetOutboundCount() *int64 {
	return s.OutboundCount
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) GetSensitiveCode() *int64 {
	return s.SensitiveCode
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) SetInfoCount(v int64) *DescribeSensitiveDetectionResultResponseBodyDataResultMax {
	s.InfoCount = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) SetOutboundCount(v int64) *DescribeSensitiveDetectionResultResponseBodyDataResultMax {
	s.OutboundCount = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) SetSensitiveCode(v int64) *DescribeSensitiveDetectionResultResponseBodyDataResultMax {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponseBodyDataResultMax) Validate() error {
	return dara.Validate(s)
}
