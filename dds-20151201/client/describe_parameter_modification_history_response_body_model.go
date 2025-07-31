// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterModificationHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHistoricalParameters(v *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) *DescribeParameterModificationHistoryResponseBody
	GetHistoricalParameters() *DescribeParameterModificationHistoryResponseBodyHistoricalParameters
	SetRequestId(v string) *DescribeParameterModificationHistoryResponseBody
	GetRequestId() *string
}

type DescribeParameterModificationHistoryResponseBody struct {
	// Details about the parameter modification records.
	HistoricalParameters *DescribeParameterModificationHistoryResponseBodyHistoricalParameters `json:"HistoricalParameters,omitempty" xml:"HistoricalParameters,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B1BB6E0E-B4EF-4145-81FA-A07719860248
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterModificationHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponseBody) GetHistoricalParameters() *DescribeParameterModificationHistoryResponseBodyHistoricalParameters {
	return s.HistoricalParameters
}

func (s *DescribeParameterModificationHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterModificationHistoryResponseBody) SetHistoricalParameters(v *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) *DescribeParameterModificationHistoryResponseBody {
	s.HistoricalParameters = v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBody) SetRequestId(v string) *DescribeParameterModificationHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterModificationHistoryResponseBodyHistoricalParameters struct {
	HistoricalParameter []*DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter `json:"HistoricalParameter,omitempty" xml:"HistoricalParameter,omitempty" type:"Repeated"`
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) GetHistoricalParameter() []*DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	return s.HistoricalParameter
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) SetHistoricalParameter(v []*DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) *DescribeParameterModificationHistoryResponseBodyHistoricalParameters {
	s.HistoricalParameter = v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter struct {
	// The time when the parameter was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-12T07:58:24Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The parameter value after modification.
	//
	// example:
	//
	// 200
	NewParameterValue *string `json:"NewParameterValue,omitempty" xml:"NewParameterValue,omitempty"`
	// The parameter value before modification.
	//
	// example:
	//
	// 100
	OldParameterValue *string `json:"OldParameterValue,omitempty" xml:"OldParameterValue,omitempty"`
	// The name of the modified parameter.
	//
	// example:
	//
	// operationProfiling.slowOpThresholdMs
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) GetNewParameterValue() *string {
	return s.NewParameterValue
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) GetOldParameterValue() *string {
	return s.OldParameterValue
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetModifyTime(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.ModifyTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetNewParameterValue(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.NewParameterValue = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetOldParameterValue(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.OldParameterValue = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) SetParameterName(v string) *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponseBodyHistoricalParametersHistoricalParameter) Validate() error {
	return dara.Validate(s)
}
