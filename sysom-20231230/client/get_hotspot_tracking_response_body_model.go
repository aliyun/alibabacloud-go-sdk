// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotTrackingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotspotTrackingResponseBody
	GetCode() *string
	SetData(v *GetHotspotTrackingResponseBodyData) *GetHotspotTrackingResponseBody
	GetData() *GetHotspotTrackingResponseBodyData
	SetMessage(v string) *GetHotspotTrackingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotTrackingResponseBody
	GetRequestId() *string
}

type GetHotspotTrackingResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotTrackingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotspotTrackingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTrackingResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotspotTrackingResponseBody) GetData() *GetHotspotTrackingResponseBodyData {
	return s.Data
}

func (s *GetHotspotTrackingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotTrackingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotTrackingResponseBody) SetCode(v string) *GetHotspotTrackingResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotTrackingResponseBody) SetData(v *GetHotspotTrackingResponseBodyData) *GetHotspotTrackingResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotTrackingResponseBody) SetMessage(v string) *GetHotspotTrackingResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotTrackingResponseBody) SetRequestId(v string) *GetHotspotTrackingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotTrackingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHotspotTrackingResponseBodyData struct {
	Flame  *GetHotspotTrackingResponseBodyDataFlame  `json:"flame,omitempty" xml:"flame,omitempty" type:"Struct"`
	Series *GetHotspotTrackingResponseBodyDataSeries `json:"series,omitempty" xml:"series,omitempty" type:"Struct"`
}

func (s GetHotspotTrackingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTrackingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBodyData) GetFlame() *GetHotspotTrackingResponseBodyDataFlame {
	return s.Flame
}

func (s *GetHotspotTrackingResponseBodyData) GetSeries() *GetHotspotTrackingResponseBodyDataSeries {
	return s.Series
}

func (s *GetHotspotTrackingResponseBodyData) SetFlame(v *GetHotspotTrackingResponseBodyDataFlame) *GetHotspotTrackingResponseBodyData {
	s.Flame = v
	return s
}

func (s *GetHotspotTrackingResponseBodyData) SetSeries(v *GetHotspotTrackingResponseBodyDataSeries) *GetHotspotTrackingResponseBodyData {
	s.Series = v
	return s
}

func (s *GetHotspotTrackingResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetHotspotTrackingResponseBodyDataFlame struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotTrackingResponseBodyDataFlame) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTrackingResponseBodyDataFlame) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBodyDataFlame) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotTrackingResponseBodyDataFlame) GetValues() [][]*string {
	return s.Values
}

func (s *GetHotspotTrackingResponseBodyDataFlame) SetColumns(v []*string) *GetHotspotTrackingResponseBodyDataFlame {
	s.Columns = v
	return s
}

func (s *GetHotspotTrackingResponseBodyDataFlame) SetValues(v [][]*string) *GetHotspotTrackingResponseBodyDataFlame {
	s.Values = v
	return s
}

func (s *GetHotspotTrackingResponseBodyDataFlame) Validate() error {
	return dara.Validate(s)
}

type GetHotspotTrackingResponseBodyDataSeries struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotTrackingResponseBodyDataSeries) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTrackingResponseBodyDataSeries) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBodyDataSeries) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotTrackingResponseBodyDataSeries) GetValues() [][]*string {
	return s.Values
}

func (s *GetHotspotTrackingResponseBodyDataSeries) SetColumns(v []*string) *GetHotspotTrackingResponseBodyDataSeries {
	s.Columns = v
	return s
}

func (s *GetHotspotTrackingResponseBodyDataSeries) SetValues(v [][]*string) *GetHotspotTrackingResponseBodyDataSeries {
	s.Values = v
	return s
}

func (s *GetHotspotTrackingResponseBodyDataSeries) Validate() error {
	return dara.Validate(s)
}
