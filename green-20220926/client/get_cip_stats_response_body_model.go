// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCipStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCipStatsResponseBody
	GetCode() *int32
	SetData(v *GetCipStatsResponseBodyData) *GetCipStatsResponseBody
	GetData() *GetCipStatsResponseBodyData
	SetHttpStatusCode(v int32) *GetCipStatsResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *GetCipStatsResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetCipStatsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCipStatsResponseBody
	GetSuccess() *bool
}

type GetCipStatsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCipStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCipStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCipStatsResponseBody) GetData() *GetCipStatsResponseBodyData {
	return s.Data
}

func (s *GetCipStatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCipStatsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetCipStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCipStatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCipStatsResponseBody) SetCode(v int32) *GetCipStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCipStatsResponseBody) SetData(v *GetCipStatsResponseBodyData) *GetCipStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBody) SetHttpStatusCode(v int32) *GetCipStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCipStatsResponseBody) SetMsg(v string) *GetCipStatsResponseBody {
	s.Msg = &v
	return s
}

func (s *GetCipStatsResponseBody) SetRequestId(v string) *GetCipStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCipStatsResponseBody) SetSuccess(v bool) *GetCipStatsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCipStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCipStatsResponseBodyData struct {
	LabelStatChart []*GetCipStatsResponseBodyDataLabelStatChart `json:"LabelStatChart,omitempty" xml:"LabelStatChart,omitempty" type:"Repeated"`
	TotalStat      map[string]map[string]interface{}            `json:"TotalStat,omitempty" xml:"TotalStat,omitempty"`
	Uids           []*string                                    `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
	X              []*string                                    `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	Y              []*GetCipStatsResponseBodyDataY              `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
	Z              []*GetCipStatsResponseBodyDataZ              `json:"Z,omitempty" xml:"Z,omitempty" type:"Repeated"`
}

func (s GetCipStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyData) GetLabelStatChart() []*GetCipStatsResponseBodyDataLabelStatChart {
	return s.LabelStatChart
}

func (s *GetCipStatsResponseBodyData) GetTotalStat() map[string]map[string]interface{} {
	return s.TotalStat
}

func (s *GetCipStatsResponseBodyData) GetUids() []*string {
	return s.Uids
}

func (s *GetCipStatsResponseBodyData) GetX() []*string {
	return s.X
}

func (s *GetCipStatsResponseBodyData) GetY() []*GetCipStatsResponseBodyDataY {
	return s.Y
}

func (s *GetCipStatsResponseBodyData) GetZ() []*GetCipStatsResponseBodyDataZ {
	return s.Z
}

func (s *GetCipStatsResponseBodyData) SetLabelStatChart(v []*GetCipStatsResponseBodyDataLabelStatChart) *GetCipStatsResponseBodyData {
	s.LabelStatChart = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetTotalStat(v map[string]map[string]interface{}) *GetCipStatsResponseBodyData {
	s.TotalStat = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetUids(v []*string) *GetCipStatsResponseBodyData {
	s.Uids = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetX(v []*string) *GetCipStatsResponseBodyData {
	s.X = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetY(v []*GetCipStatsResponseBodyDataY) *GetCipStatsResponseBodyData {
	s.Y = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetZ(v []*GetCipStatsResponseBodyDataZ) *GetCipStatsResponseBodyData {
	s.Z = v
	return s
}

func (s *GetCipStatsResponseBodyData) Validate() error {
	if s.LabelStatChart != nil {
		for _, item := range s.LabelStatChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Y != nil {
		for _, item := range s.Y {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Z != nil {
		for _, item := range s.Z {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCipStatsResponseBodyDataLabelStatChart struct {
	ImageTreeChar []*GetCipStatsResponseBodyDataLabelStatChartImageTreeChar `json:"ImageTreeChar,omitempty" xml:"ImageTreeChar,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	ServiceCode   *string                                                   `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	TextTreeChart []*GetCipStatsResponseBodyDataLabelStatChartTextTreeChart `json:"TextTreeChart,omitempty" xml:"TextTreeChart,omitempty" type:"Repeated"`
	// example:
	//
	// 117
	TotalCount     *int64                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TreeChart      []*GetCipStatsResponseBodyDataLabelStatChartTreeChart      `json:"TreeChart,omitempty" xml:"TreeChart,omitempty" type:"Repeated"`
	VoiceTreeChart []*GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart `json:"VoiceTreeChart,omitempty" xml:"VoiceTreeChart,omitempty" type:"Repeated"`
	X              []*string                                                  `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	Y              []*GetCipStatsResponseBodyDataLabelStatChartY              `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s GetCipStatsResponseBodyDataLabelStatChart) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetImageTreeChar() []*GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	return s.ImageTreeChar
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetTextTreeChart() []*GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	return s.TextTreeChart
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetTreeChart() []*GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	return s.TreeChart
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetVoiceTreeChart() []*GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	return s.VoiceTreeChart
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetX() []*string {
	return s.X
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) GetY() []*GetCipStatsResponseBodyDataLabelStatChartY {
	return s.Y
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetImageTreeChar(v []*GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) *GetCipStatsResponseBodyDataLabelStatChart {
	s.ImageTreeChar = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetServiceCode(v string) *GetCipStatsResponseBodyDataLabelStatChart {
	s.ServiceCode = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetTextTreeChart(v []*GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) *GetCipStatsResponseBodyDataLabelStatChart {
	s.TextTreeChart = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetTotalCount(v int64) *GetCipStatsResponseBodyDataLabelStatChart {
	s.TotalCount = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetTreeChart(v []*GetCipStatsResponseBodyDataLabelStatChartTreeChart) *GetCipStatsResponseBodyDataLabelStatChart {
	s.TreeChart = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetVoiceTreeChart(v []*GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) *GetCipStatsResponseBodyDataLabelStatChart {
	s.VoiceTreeChart = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetX(v []*string) *GetCipStatsResponseBodyDataLabelStatChart {
	s.X = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetY(v []*GetCipStatsResponseBodyDataLabelStatChartY) *GetCipStatsResponseBodyDataLabelStatChart {
	s.Y = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) Validate() error {
	if s.ImageTreeChar != nil {
		for _, item := range s.ImageTreeChar {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextTreeChart != nil {
		for _, item := range s.TextTreeChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TreeChart != nil {
		for _, item := range s.TreeChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VoiceTreeChart != nil {
		for _, item := range s.VoiceTreeChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Y != nil {
		for _, item := range s.Y {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCipStatsResponseBodyDataLabelStatChartImageTreeChar struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) GetDescription() *string {
	return s.Description
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) GetValue() *string {
	return s.Value
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	s.Value = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) Validate() error {
	return dara.Validate(s)
}

type GetCipStatsResponseBodyDataLabelStatChartTextTreeChart struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) GetDescription() *string {
	return s.Description
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) GetValue() *string {
	return s.Value
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	s.Value = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) Validate() error {
	return dara.Validate(s)
}

type GetCipStatsResponseBodyDataLabelStatChartTreeChart struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 99.91
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartTreeChart) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartTreeChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) GetDescription() *string {
	return s.Description
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) GetValue() *string {
	return s.Value
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	s.Value = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) Validate() error {
	return dara.Validate(s)
}

type GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 99.91
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) GetDescription() *string {
	return s.Description
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) GetValue() *string {
	return s.Value
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	s.Value = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) Validate() error {
	return dara.Validate(s)
}

type GetCipStatsResponseBodyDataLabelStatChartY struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartY) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartY) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) GetData() []*int64 {
	return s.Data
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) SetData(v []*int64) *GetCipStatsResponseBodyDataLabelStatChartY {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartY {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) Validate() error {
	return dara.Validate(s)
}

type GetCipStatsResponseBodyDataY struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCipStatsResponseBodyDataY) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataY) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataY) GetData() []*int64 {
	return s.Data
}

func (s *GetCipStatsResponseBodyDataY) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataY) SetData(v []*int64) *GetCipStatsResponseBodyDataY {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBodyDataY) SetName(v string) *GetCipStatsResponseBodyDataY {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataY) Validate() error {
	return dara.Validate(s)
}

type GetCipStatsResponseBodyDataZ struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCipStatsResponseBodyDataZ) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsResponseBodyDataZ) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataZ) GetData() []*int64 {
	return s.Data
}

func (s *GetCipStatsResponseBodyDataZ) GetName() *string {
	return s.Name
}

func (s *GetCipStatsResponseBodyDataZ) SetData(v []*int64) *GetCipStatsResponseBodyDataZ {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBodyDataZ) SetName(v string) *GetCipStatsResponseBodyDataZ {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataZ) Validate() error {
	return dara.Validate(s)
}
