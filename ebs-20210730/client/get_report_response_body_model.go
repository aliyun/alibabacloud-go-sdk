// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatas(v []*GetReportResponseBodyDatas) *GetReportResponseBody
	GetDatas() []*GetReportResponseBodyDatas
	SetRequestId(v string) *GetReportResponseBody
	GetRequestId() *string
}

type GetReportResponseBody struct {
	// Data Details.
	Datas []*GetReportResponseBodyDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportResponseBody) GetDatas() []*GetReportResponseBodyDatas {
	return s.Datas
}

func (s *GetReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReportResponseBody) SetDatas(v []*GetReportResponseBodyDatas) *GetReportResponseBody {
	s.Datas = v
	return s
}

func (s *GetReportResponseBody) SetRequestId(v string) *GetReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportResponseBody) Validate() error {
	if s.Datas != nil {
		for _, item := range s.Datas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetReportResponseBodyDatas struct {
	// Data.
	Data []*GetReportResponseBodyDatasData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Data Title.
	//
	// example:
	//
	// disk_count_percent_by_category
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetReportResponseBodyDatas) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *GetReportResponseBodyDatas) GetData() []*GetReportResponseBodyDatasData {
	return s.Data
}

func (s *GetReportResponseBodyDatas) GetTitle() *string {
	return s.Title
}

func (s *GetReportResponseBodyDatas) SetData(v []*GetReportResponseBodyDatasData) *GetReportResponseBodyDatas {
	s.Data = v
	return s
}

func (s *GetReportResponseBodyDatas) SetTitle(v string) *GetReportResponseBodyDatas {
	s.Title = &v
	return s
}

func (s *GetReportResponseBodyDatas) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetReportResponseBodyDatasData struct {
	// Data Points.
	//
	// example:
	//
	// {
	//
	//   "1726416000": 0.44,
	//
	//   "1726502400": 0.44,
	//
	//   "1726588800": 0.44,
	//
	//   "1726675200": 0.44,
	//
	//   "1726761600": 0.43,
	//
	//   "1726848000": 0.43,
	//
	//   "1726934400": 0.43,
	//
	//   "1727020800": 0.43
	//
	// }
	DataPoints map[string]interface{} `json:"DataPoints,omitempty" xml:"DataPoints,omitempty"`
	// Data Labels.
	//
	// example:
	//
	// {
	//
	//   "category": "cloud"
	//
	// }
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s GetReportResponseBodyDatasData) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseBodyDatasData) GoString() string {
	return s.String()
}

func (s *GetReportResponseBodyDatasData) GetDataPoints() map[string]interface{} {
	return s.DataPoints
}

func (s *GetReportResponseBodyDatasData) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *GetReportResponseBodyDatasData) SetDataPoints(v map[string]interface{}) *GetReportResponseBodyDatasData {
	s.DataPoints = v
	return s
}

func (s *GetReportResponseBodyDatasData) SetLabels(v map[string]interface{}) *GetReportResponseBodyDatasData {
	s.Labels = v
	return s
}

func (s *GetReportResponseBodyDatasData) Validate() error {
	return dara.Validate(s)
}
