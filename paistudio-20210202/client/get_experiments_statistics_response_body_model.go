// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentsStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetExperimentsStatisticsResponseBodyData) *GetExperimentsStatisticsResponseBody
	GetData() []*GetExperimentsStatisticsResponseBodyData
	SetRequestId(v string) *GetExperimentsStatisticsResponseBody
	GetRequestId() *string
}

type GetExperimentsStatisticsResponseBody struct {
	Data []*GetExperimentsStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExperimentsStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsResponseBody) GetData() []*GetExperimentsStatisticsResponseBodyData {
	return s.Data
}

func (s *GetExperimentsStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentsStatisticsResponseBody) SetData(v []*GetExperimentsStatisticsResponseBodyData) *GetExperimentsStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetExperimentsStatisticsResponseBody) SetRequestId(v string) *GetExperimentsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentsStatisticsResponseBody) Validate() error {
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

type GetExperimentsStatisticsResponseBodyData struct {
	// example:
	//
	// 10
	CreateCount *int64 `json:"CreateCount,omitempty" xml:"CreateCount,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 16381
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentsStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentsStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsResponseBodyData) GetCreateCount() *int64 {
	return s.CreateCount
}

func (s *GetExperimentsStatisticsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetExperimentsStatisticsResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExperimentsStatisticsResponseBodyData) SetCreateCount(v int64) *GetExperimentsStatisticsResponseBodyData {
	s.CreateCount = &v
	return s
}

func (s *GetExperimentsStatisticsResponseBodyData) SetTotalCount(v int64) *GetExperimentsStatisticsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetExperimentsStatisticsResponseBodyData) SetWorkspaceId(v string) *GetExperimentsStatisticsResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetExperimentsStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
