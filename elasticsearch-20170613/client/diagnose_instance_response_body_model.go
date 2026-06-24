// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DiagnoseInstanceResponseBody
	GetRequestId() *string
	SetResult(v *DiagnoseInstanceResponseBodyResult) *DiagnoseInstanceResponseBody
	GetResult() *DiagnoseInstanceResponseBodyResult
}

type DiagnoseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DiagnoseInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DiagnoseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiagnoseInstanceResponseBody) GetResult() *DiagnoseInstanceResponseBodyResult {
	return s.Result
}

func (s *DiagnoseInstanceResponseBody) SetRequestId(v string) *DiagnoseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnoseInstanceResponseBody) SetResult(v *DiagnoseInstanceResponseBodyResult) *DiagnoseInstanceResponseBody {
	s.Result = v
	return s
}

func (s *DiagnoseInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiagnoseInstanceResponseBodyResult struct {
	// The timestamp when the diagnostic report was generated.
	//
	// example:
	//
	// 1535745731000
	CreateTime    *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiagnosisMode *string `json:"diagnosisMode,omitempty" xml:"diagnosisMode,omitempty"`
	// The instance ID of the diagnosed instance.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****
	InstanceId *string                                    `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Items      []*DiagnoseInstanceResponseBodyResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The report ID.
	//
	// example:
	//
	// trigger__2020-08-17T17:09:02
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The diagnostic status. Valid values: SUCCESS, FAILED, and RUNNING.
	//
	// example:
	//
	// RUNNING
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DiagnoseInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DiagnoseInstanceResponseBodyResult) GetDiagnosisMode() *string {
	return s.DiagnosisMode
}

func (s *DiagnoseInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DiagnoseInstanceResponseBodyResult) GetItems() []*DiagnoseInstanceResponseBodyResultItems {
	return s.Items
}

func (s *DiagnoseInstanceResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *DiagnoseInstanceResponseBodyResult) GetState() *string {
	return s.State
}

func (s *DiagnoseInstanceResponseBodyResult) SetCreateTime(v int64) *DiagnoseInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetDiagnosisMode(v string) *DiagnoseInstanceResponseBodyResult {
	s.DiagnosisMode = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetInstanceId(v string) *DiagnoseInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetItems(v []*DiagnoseInstanceResponseBodyResultItems) *DiagnoseInstanceResponseBodyResult {
	s.Items = v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetReportId(v string) *DiagnoseInstanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetState(v string) *DiagnoseInstanceResponseBodyResult {
	s.State = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DiagnoseInstanceResponseBodyResultItems struct {
	Desc    *string                `json:"desc,omitempty" xml:"desc,omitempty"`
	Detail  map[string]interface{} `json:"detail,omitempty" xml:"detail,omitempty"`
	Item    *string                `json:"item,omitempty" xml:"item,omitempty"`
	Name    *string                `json:"name,omitempty" xml:"name,omitempty"`
	State   *string                `json:"state,omitempty" xml:"state,omitempty"`
	Suggest *string                `json:"suggest,omitempty" xml:"suggest,omitempty"`
}

func (s DiagnoseInstanceResponseBodyResultItems) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBodyResultItems) GetDesc() *string {
	return s.Desc
}

func (s *DiagnoseInstanceResponseBodyResultItems) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *DiagnoseInstanceResponseBodyResultItems) GetItem() *string {
	return s.Item
}

func (s *DiagnoseInstanceResponseBodyResultItems) GetName() *string {
	return s.Name
}

func (s *DiagnoseInstanceResponseBodyResultItems) GetState() *string {
	return s.State
}

func (s *DiagnoseInstanceResponseBodyResultItems) GetSuggest() *string {
	return s.Suggest
}

func (s *DiagnoseInstanceResponseBodyResultItems) SetDesc(v string) *DiagnoseInstanceResponseBodyResultItems {
	s.Desc = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResultItems) SetDetail(v map[string]interface{}) *DiagnoseInstanceResponseBodyResultItems {
	s.Detail = v
	return s
}

func (s *DiagnoseInstanceResponseBodyResultItems) SetItem(v string) *DiagnoseInstanceResponseBodyResultItems {
	s.Item = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResultItems) SetName(v string) *DiagnoseInstanceResponseBodyResultItems {
	s.Name = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResultItems) SetState(v string) *DiagnoseInstanceResponseBodyResultItems {
	s.State = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResultItems) SetSuggest(v string) *DiagnoseInstanceResponseBodyResultItems {
	s.Suggest = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResultItems) Validate() error {
	return dara.Validate(s)
}
