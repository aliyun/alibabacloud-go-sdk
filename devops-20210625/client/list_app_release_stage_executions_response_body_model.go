// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppReleaseStageExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListAppReleaseStageExecutionsResponseBody
	GetCurrent() *int64
	SetData(v []*ListAppReleaseStageExecutionsResponseBodyData) *ListAppReleaseStageExecutionsResponseBody
	GetData() []*ListAppReleaseStageExecutionsResponseBodyData
	SetNextToken(v string) *ListAppReleaseStageExecutionsResponseBody
	GetNextToken() *string
	SetPages(v int64) *ListAppReleaseStageExecutionsResponseBody
	GetPages() *int64
	SetPerPage(v int64) *ListAppReleaseStageExecutionsResponseBody
	GetPerPage() *int64
	SetTotal(v int64) *ListAppReleaseStageExecutionsResponseBody
	GetTotal() *int64
}

type ListAppReleaseStageExecutionsResponseBody struct {
	// example:
	//
	// 1
	Current *int64                                           `json:"current,omitempty" xml:"current,omitempty"`
	Data    []*ListAppReleaseStageExecutionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// vxc2341gfssad12
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// example:
	//
	// 20
	PerPage *int64 `json:"perPage,omitempty" xml:"perPage,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAppReleaseStageExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionsResponseBody) GetCurrent() *int64 {
	return s.Current
}

func (s *ListAppReleaseStageExecutionsResponseBody) GetData() []*ListAppReleaseStageExecutionsResponseBodyData {
	return s.Data
}

func (s *ListAppReleaseStageExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppReleaseStageExecutionsResponseBody) GetPages() *int64 {
	return s.Pages
}

func (s *ListAppReleaseStageExecutionsResponseBody) GetPerPage() *int64 {
	return s.PerPage
}

func (s *ListAppReleaseStageExecutionsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAppReleaseStageExecutionsResponseBody) SetCurrent(v int64) *ListAppReleaseStageExecutionsResponseBody {
	s.Current = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBody) SetData(v []*ListAppReleaseStageExecutionsResponseBodyData) *ListAppReleaseStageExecutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBody) SetNextToken(v string) *ListAppReleaseStageExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBody) SetPages(v int64) *ListAppReleaseStageExecutionsResponseBody {
	s.Pages = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBody) SetPerPage(v int64) *ListAppReleaseStageExecutionsResponseBody {
	s.PerPage = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBody) SetTotal(v int64) *ListAppReleaseStageExecutionsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBody) Validate() error {
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

type ListAppReleaseStageExecutionsResponseBodyData struct {
	// example:
	//
	// 2024-06-25T07:26:18.000+00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// 2024-06-25T07:25:54.000+00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// MANUAL
	TriggerMode *string `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
}

func (s ListAppReleaseStageExecutionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) SetEndTime(v string) *ListAppReleaseStageExecutionsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) SetNumber(v string) *ListAppReleaseStageExecutionsResponseBodyData {
	s.Number = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) SetStartTime(v string) *ListAppReleaseStageExecutionsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) SetState(v string) *ListAppReleaseStageExecutionsResponseBodyData {
	s.State = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) SetTriggerMode(v string) *ListAppReleaseStageExecutionsResponseBodyData {
	s.TriggerMode = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
