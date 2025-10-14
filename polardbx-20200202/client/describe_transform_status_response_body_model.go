// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransformStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeTransformStatusResponseBodyData) *DescribeTransformStatusResponseBody
	GetData() *DescribeTransformStatusResponseBodyData
	SetRequestId(v string) *DescribeTransformStatusResponseBody
	GetRequestId() *string
}

type DescribeTransformStatusResponseBody struct {
	Data *DescribeTransformStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTransformStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransformStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransformStatusResponseBody) GetData() *DescribeTransformStatusResponseBodyData {
	return s.Data
}

func (s *DescribeTransformStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransformStatusResponseBody) SetData(v *DescribeTransformStatusResponseBodyData) *DescribeTransformStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTransformStatusResponseBody) SetRequestId(v string) *DescribeTransformStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransformStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTransformStatusResponseBodyData struct {
	// example:
	//
	// true
	CanCancel *bool `json:"CanCancel,omitempty" xml:"CanCancel,omitempty"`
	// example:
	//
	// true
	CanFinish     *bool `json:"CanFinish,omitempty" xml:"CanFinish,omitempty"`
	CanSwitch     *bool `json:"CanSwitch,omitempty" xml:"CanSwitch,omitempty"`
	CanUndoSwitch *bool `json:"CanUndoSwitch,omitempty" xml:"CanUndoSwitch,omitempty"`
	// example:
	//
	// pxc-*****
	EnterpriseInsName *string `json:"EnterpriseInsName,omitempty" xml:"EnterpriseInsName,omitempty"`
	// example:
	//
	// none
	Phase         *string                `json:"Phase,omitempty" xml:"Phase,omitempty"`
	ReportSummary map[string]interface{} `json:"ReportSummary,omitempty" xml:"ReportSummary,omitempty"`
	// example:
	//
	// 2025-09-02 16:52:47.0
	ReportTime *int64 `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	// example:
	//
	// pxc-*****
	StandardInsName *string `json:"StandardInsName,omitempty" xml:"StandardInsName,omitempty"`
}

func (s DescribeTransformStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransformStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTransformStatusResponseBodyData) GetCanCancel() *bool {
	return s.CanCancel
}

func (s *DescribeTransformStatusResponseBodyData) GetCanFinish() *bool {
	return s.CanFinish
}

func (s *DescribeTransformStatusResponseBodyData) GetCanSwitch() *bool {
	return s.CanSwitch
}

func (s *DescribeTransformStatusResponseBodyData) GetCanUndoSwitch() *bool {
	return s.CanUndoSwitch
}

func (s *DescribeTransformStatusResponseBodyData) GetEnterpriseInsName() *string {
	return s.EnterpriseInsName
}

func (s *DescribeTransformStatusResponseBodyData) GetPhase() *string {
	return s.Phase
}

func (s *DescribeTransformStatusResponseBodyData) GetReportSummary() map[string]interface{} {
	return s.ReportSummary
}

func (s *DescribeTransformStatusResponseBodyData) GetReportTime() *int64 {
	return s.ReportTime
}

func (s *DescribeTransformStatusResponseBodyData) GetStandardInsName() *string {
	return s.StandardInsName
}

func (s *DescribeTransformStatusResponseBodyData) SetCanCancel(v bool) *DescribeTransformStatusResponseBodyData {
	s.CanCancel = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetCanFinish(v bool) *DescribeTransformStatusResponseBodyData {
	s.CanFinish = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetCanSwitch(v bool) *DescribeTransformStatusResponseBodyData {
	s.CanSwitch = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetCanUndoSwitch(v bool) *DescribeTransformStatusResponseBodyData {
	s.CanUndoSwitch = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetEnterpriseInsName(v string) *DescribeTransformStatusResponseBodyData {
	s.EnterpriseInsName = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetPhase(v string) *DescribeTransformStatusResponseBodyData {
	s.Phase = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetReportSummary(v map[string]interface{}) *DescribeTransformStatusResponseBodyData {
	s.ReportSummary = v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetReportTime(v int64) *DescribeTransformStatusResponseBodyData {
	s.ReportTime = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) SetStandardInsName(v string) *DescribeTransformStatusResponseBodyData {
	s.StandardInsName = &v
	return s
}

func (s *DescribeTransformStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
