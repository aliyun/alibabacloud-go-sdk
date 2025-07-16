// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeParameterTemplatesResponseBodyData) *DescribeParameterTemplatesResponseBody
	GetData() *DescribeParameterTemplatesResponseBodyData
	SetRequestId(v string) *DescribeParameterTemplatesResponseBody
	GetRequestId() *string
}

type DescribeParameterTemplatesResponseBody struct {
	Data *DescribeParameterTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBody) GetData() *DescribeParameterTemplatesResponseBodyData {
	return s.Data
}

func (s *DescribeParameterTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterTemplatesResponseBody) SetData(v *DescribeParameterTemplatesResponseBodyData) *DescribeParameterTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetRequestId(v string) *DescribeParameterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterTemplatesResponseBodyData struct {
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 218
	ParameterCount *int32                                                  `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	Parameters     []*DescribeParameterTemplatesResponseBodyDataParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DescribeParameterTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *DescribeParameterTemplatesResponseBodyData) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeParameterTemplatesResponseBodyData) GetParameterCount() *int32 {
	return s.ParameterCount
}

func (s *DescribeParameterTemplatesResponseBodyData) GetParameters() []*DescribeParameterTemplatesResponseBodyDataParameters {
	return s.Parameters
}

func (s *DescribeParameterTemplatesResponseBodyData) SetEngine(v string) *DescribeParameterTemplatesResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) SetEngineVersion(v string) *DescribeParameterTemplatesResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) SetParameterCount(v int32) *DescribeParameterTemplatesResponseBodyData {
	s.ParameterCount = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) SetParameters(v []*DescribeParameterTemplatesResponseBodyDataParameters) *DescribeParameterTemplatesResponseBodyData {
	s.Parameters = v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeParameterTemplatesResponseBodyDataParameters struct {
	// example:
	//
	// [0|1]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// example:
	//
	// 0
	Dynamic *int32 `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// example:
	//
	// polarx hatp addition param
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// example:
	//
	// loose_enable_gts
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	// example:
	//
	// 0
	Revisable *int32 `json:"Revisable,omitempty" xml:"Revisable,omitempty"`
}

func (s DescribeParameterTemplatesResponseBodyDataParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyDataParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) GetDynamic() *int32 {
	return s.Dynamic
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) GetRevisable() *int32 {
	return s.Revisable
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetCheckingCode(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetDynamic(v int32) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.Dynamic = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetParameterDescription(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetParameterName(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetParameterValue(v string) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) SetRevisable(v int32) *DescribeParameterTemplatesResponseBodyDataParameters {
	s.Revisable = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyDataParameters) Validate() error {
	return dara.Validate(s)
}
