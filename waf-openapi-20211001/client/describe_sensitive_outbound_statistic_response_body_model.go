// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveOutboundStatisticResponseBodyData) *DescribeSensitiveOutboundStatisticResponseBody
	GetData() []*DescribeSensitiveOutboundStatisticResponseBodyData
	SetRequestId(v string) *DescribeSensitiveOutboundStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveOutboundStatisticResponseBody
	GetTotalCount() *int64
}

type DescribeSensitiveOutboundStatisticResponseBody struct {
	// The data types of personal information involved in cross-border data transfer.
	Data []*DescribeSensitiveOutboundStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSensitiveOutboundStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) GetData() []*DescribeSensitiveOutboundStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) SetData(v []*DescribeSensitiveOutboundStatisticResponseBodyData) *DescribeSensitiveOutboundStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) SetRequestId(v string) *DescribeSensitiveOutboundStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) SetTotalCount(v int64) *DescribeSensitiveOutboundStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSensitiveOutboundStatisticResponseBodyData struct {
	// The evaluation result. Valid values:
	//
	// 	- **report**: Risks exist in cross-border data transfer.
	//
	// 	- **none**: No risks exist in cross-border data transfer.
	//
	// example:
	//
	// report
	DetectionResult *string `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 546
	InfoCount *int64 `json:"InfoCount,omitempty" xml:"InfoCount,omitempty"`
	// The number of data entries that are transferred across borders.
	//
	// example:
	//
	// 300
	OutboundCount *int64 `json:"OutboundCount,omitempty" xml:"OutboundCount,omitempty"`
	// The type of the sensitive data.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported types of sensitive data.
	//
	// example:
	//
	// 1001
	SensitiveCode *int64 `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	// The sensitivity level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The type of the information. Valid values:
	//
	// 	- **info**: full personal information
	//
	// 	- **sensitive**: sensitive personal information
	//
	// example:
	//
	// info
	SensitiveType *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
}

func (s DescribeSensitiveOutboundStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) GetDetectionResult() *string {
	return s.DetectionResult
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) GetInfoCount() *int64 {
	return s.InfoCount
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) GetOutboundCount() *int64 {
	return s.OutboundCount
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) GetSensitiveCode() *int64 {
	return s.SensitiveCode
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) GetSensitiveType() *string {
	return s.SensitiveType
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) SetDetectionResult(v string) *DescribeSensitiveOutboundStatisticResponseBodyData {
	s.DetectionResult = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) SetInfoCount(v int64) *DescribeSensitiveOutboundStatisticResponseBodyData {
	s.InfoCount = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) SetOutboundCount(v int64) *DescribeSensitiveOutboundStatisticResponseBodyData {
	s.OutboundCount = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) SetSensitiveCode(v int64) *DescribeSensitiveOutboundStatisticResponseBodyData {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) SetSensitiveLevel(v string) *DescribeSensitiveOutboundStatisticResponseBodyData {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) SetSensitiveType(v string) *DescribeSensitiveOutboundStatisticResponseBodyData {
	s.SensitiveType = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
