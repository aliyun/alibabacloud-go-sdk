// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFixCycleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeImageFixCycleConfigResponseBodyData) *DescribeImageFixCycleConfigResponseBody
	GetData() *DescribeImageFixCycleConfigResponseBodyData
	SetRequestId(v string) *DescribeImageFixCycleConfigResponseBody
	GetRequestId() *string
}

type DescribeImageFixCycleConfigResponseBody struct {
	// The response parameters.
	Data *DescribeImageFixCycleConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageFixCycleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixCycleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageFixCycleConfigResponseBody) GetData() *DescribeImageFixCycleConfigResponseBodyData {
	return s.Data
}

func (s *DescribeImageFixCycleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageFixCycleConfigResponseBody) SetData(v *DescribeImageFixCycleConfigResponseBodyData) *DescribeImageFixCycleConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageFixCycleConfigResponseBody) SetRequestId(v string) *DescribeImageFixCycleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageFixCycleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageFixCycleConfigResponseBodyData struct {
	// The cycle of the scheduled fix. Unit: day.
	//
	// example:
	//
	// 7
	ImageFixCycle *int32 `json:"ImageFixCycle,omitempty" xml:"ImageFixCycle,omitempty"`
	// Indicates whether the scheduled fix of image risks is enabled.
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
	//
	// example:
	//
	// on
	ImageFixSwitch *string `json:"ImageFixSwitch,omitempty" xml:"ImageFixSwitch,omitempty"`
	// The range of the scheduled fix. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: The type of the image risk. The value is fixed to repo.
	//
	// 	- **target**: The content of the image risk. The value is in the format of Namespace/Image repository.
	//
	// example:
	//
	// {\\"type\\":\\"repo\\",\\"target\\":[\\"qa-dac/yyuan9\\",\\"cdp-uat/zentao\\",\\"cafdms-qa/xxl-job-admin\\",\\"cafdms-qa/utils/jdk\\",\\"cafmfbi/ui\\",\\"cdp-uat/tradingdesk-webapp\\"]}
	ImageFixTarget *string `json:"ImageFixTarget,omitempty" xml:"ImageFixTarget,omitempty"`
	// The time range during which the image was modified. Unit: day.
	//
	// example:
	//
	// 30
	ImageTimeRange *int32 `json:"ImageTimeRange,omitempty" xml:"ImageTimeRange,omitempty"`
}

func (s DescribeImageFixCycleConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixCycleConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageFixCycleConfigResponseBodyData) GetImageFixCycle() *int32 {
	return s.ImageFixCycle
}

func (s *DescribeImageFixCycleConfigResponseBodyData) GetImageFixSwitch() *string {
	return s.ImageFixSwitch
}

func (s *DescribeImageFixCycleConfigResponseBodyData) GetImageFixTarget() *string {
	return s.ImageFixTarget
}

func (s *DescribeImageFixCycleConfigResponseBodyData) GetImageTimeRange() *int32 {
	return s.ImageTimeRange
}

func (s *DescribeImageFixCycleConfigResponseBodyData) SetImageFixCycle(v int32) *DescribeImageFixCycleConfigResponseBodyData {
	s.ImageFixCycle = &v
	return s
}

func (s *DescribeImageFixCycleConfigResponseBodyData) SetImageFixSwitch(v string) *DescribeImageFixCycleConfigResponseBodyData {
	s.ImageFixSwitch = &v
	return s
}

func (s *DescribeImageFixCycleConfigResponseBodyData) SetImageFixTarget(v string) *DescribeImageFixCycleConfigResponseBodyData {
	s.ImageFixTarget = &v
	return s
}

func (s *DescribeImageFixCycleConfigResponseBodyData) SetImageTimeRange(v int32) *DescribeImageFixCycleConfigResponseBodyData {
	s.ImageTimeRange = &v
	return s
}

func (s *DescribeImageFixCycleConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
