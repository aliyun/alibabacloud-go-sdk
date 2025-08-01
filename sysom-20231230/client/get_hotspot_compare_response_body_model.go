// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotCompareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotspotCompareResponseBody
	GetCode() *string
	SetData(v *GetHotspotCompareResponseBodyData) *GetHotspotCompareResponseBody
	GetData() *GetHotspotCompareResponseBodyData
	SetMessage(v string) *GetHotspotCompareResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotCompareResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotspotCompareResponseBody
	GetSuccess() *bool
}

type GetHotspotCompareResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotCompareResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotspotCompareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotspotCompareResponseBody) GetData() *GetHotspotCompareResponseBodyData {
	return s.Data
}

func (s *GetHotspotCompareResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotCompareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotCompareResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotspotCompareResponseBody) SetCode(v string) *GetHotspotCompareResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotCompareResponseBody) SetData(v *GetHotspotCompareResponseBodyData) *GetHotspotCompareResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotCompareResponseBody) SetMessage(v string) *GetHotspotCompareResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotCompareResponseBody) SetRequestId(v string) *GetHotspotCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotCompareResponseBody) SetSuccess(v bool) *GetHotspotCompareResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotCompareResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHotspotCompareResponseBodyData struct {
	Flame           *GetHotspotCompareResponseBodyDataFlame           `json:"flame,omitempty" xml:"flame,omitempty" type:"Struct"`
	SeriesInstance1 *GetHotspotCompareResponseBodyDataSeriesInstance1 `json:"series_instance1,omitempty" xml:"series_instance1,omitempty" type:"Struct"`
	SeriesInstance2 *GetHotspotCompareResponseBodyDataSeriesInstance2 `json:"series_instance2,omitempty" xml:"series_instance2,omitempty" type:"Struct"`
}

func (s GetHotspotCompareResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyData) GetFlame() *GetHotspotCompareResponseBodyDataFlame {
	return s.Flame
}

func (s *GetHotspotCompareResponseBodyData) GetSeriesInstance1() *GetHotspotCompareResponseBodyDataSeriesInstance1 {
	return s.SeriesInstance1
}

func (s *GetHotspotCompareResponseBodyData) GetSeriesInstance2() *GetHotspotCompareResponseBodyDataSeriesInstance2 {
	return s.SeriesInstance2
}

func (s *GetHotspotCompareResponseBodyData) SetFlame(v *GetHotspotCompareResponseBodyDataFlame) *GetHotspotCompareResponseBodyData {
	s.Flame = v
	return s
}

func (s *GetHotspotCompareResponseBodyData) SetSeriesInstance1(v *GetHotspotCompareResponseBodyDataSeriesInstance1) *GetHotspotCompareResponseBodyData {
	s.SeriesInstance1 = v
	return s
}

func (s *GetHotspotCompareResponseBodyData) SetSeriesInstance2(v *GetHotspotCompareResponseBodyDataSeriesInstance2) *GetHotspotCompareResponseBodyData {
	s.SeriesInstance2 = v
	return s
}

func (s *GetHotspotCompareResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetHotspotCompareResponseBodyDataFlame struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotCompareResponseBodyDataFlame) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareResponseBodyDataFlame) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyDataFlame) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotCompareResponseBodyDataFlame) GetValues() [][]*string {
	return s.Values
}

func (s *GetHotspotCompareResponseBodyDataFlame) SetColumns(v []*string) *GetHotspotCompareResponseBodyDataFlame {
	s.Columns = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataFlame) SetValues(v [][]*string) *GetHotspotCompareResponseBodyDataFlame {
	s.Values = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataFlame) Validate() error {
	return dara.Validate(s)
}

type GetHotspotCompareResponseBodyDataSeriesInstance1 struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance1) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance1) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) GetValues() [][]*string {
	return s.Values
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) SetColumns(v []*string) *GetHotspotCompareResponseBodyDataSeriesInstance1 {
	s.Columns = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) SetValues(v [][]*string) *GetHotspotCompareResponseBodyDataSeriesInstance1 {
	s.Values = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) Validate() error {
	return dara.Validate(s)
}

type GetHotspotCompareResponseBodyDataSeriesInstance2 struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance2) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance2) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) GetValues() [][]*string {
	return s.Values
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) SetColumns(v []*string) *GetHotspotCompareResponseBodyDataSeriesInstance2 {
	s.Columns = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) SetValues(v [][]*string) *GetHotspotCompareResponseBodyDataSeriesInstance2 {
	s.Values = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) Validate() error {
	return dara.Validate(s)
}
