// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentDashboardsResponseBody
	GetCode() *int32
	SetData(v *ListEnvironmentDashboardsResponseBodyData) *ListEnvironmentDashboardsResponseBody
	GetData() *ListEnvironmentDashboardsResponseBodyData
	SetMessage(v string) *ListEnvironmentDashboardsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentDashboardsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentDashboardsResponseBody
	GetSuccess() *bool
}

type ListEnvironmentDashboardsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	Data *ListEnvironmentDashboardsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C647A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDashboardsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentDashboardsResponseBody) GetData() *ListEnvironmentDashboardsResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentDashboardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentDashboardsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentDashboardsResponseBody) SetCode(v int32) *ListEnvironmentDashboardsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBody) SetData(v *ListEnvironmentDashboardsResponseBodyData) *ListEnvironmentDashboardsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentDashboardsResponseBody) SetMessage(v string) *ListEnvironmentDashboardsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBody) SetRequestId(v string) *ListEnvironmentDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBody) SetSuccess(v bool) *ListEnvironmentDashboardsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEnvironmentDashboardsResponseBodyData struct {
	// The dashboards.
	Dashboards []*ListEnvironmentDashboardsResponseBodyDataDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEnvironmentDashboardsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentDashboardsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDashboardsResponseBodyData) GetDashboards() []*ListEnvironmentDashboardsResponseBodyDataDashboards {
	return s.Dashboards
}

func (s *ListEnvironmentDashboardsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListEnvironmentDashboardsResponseBodyData) SetDashboards(v []*ListEnvironmentDashboardsResponseBodyDataDashboards) *ListEnvironmentDashboardsResponseBodyData {
	s.Dashboards = v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyData) SetTotal(v int64) *ListEnvironmentDashboardsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyData) Validate() error {
	if s.Dashboards != nil {
		for _, item := range s.Dashboards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvironmentDashboardsResponseBodyDataDashboards struct {
	// The UID of the folder.
	//
	// example:
	//
	// 1374923841627893
	FolderUid *string `json:"FolderUid,omitempty" xml:"FolderUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The keyword.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The title of the Grafana dashboard.
	//
	// example:
	//
	// kafka-instance
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The unique identifier of the dashboard.
	//
	// example:
	//
	// 1537863211936042
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The complete URL of the dashboard.
	//
	// example:
	//
	// http://xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListEnvironmentDashboardsResponseBodyDataDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentDashboardsResponseBodyDataDashboards) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) GetFolderUid() *string {
	return s.FolderUid
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) GetRegion() *string {
	return s.Region
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) GetTags() []*string {
	return s.Tags
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) GetTitle() *string {
	return s.Title
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) GetUid() *string {
	return s.Uid
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) SetFolderUid(v string) *ListEnvironmentDashboardsResponseBodyDataDashboards {
	s.FolderUid = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) SetRegion(v string) *ListEnvironmentDashboardsResponseBodyDataDashboards {
	s.Region = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) SetTags(v []*string) *ListEnvironmentDashboardsResponseBodyDataDashboards {
	s.Tags = v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) SetTitle(v string) *ListEnvironmentDashboardsResponseBodyDataDashboards {
	s.Title = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) SetUid(v string) *ListEnvironmentDashboardsResponseBodyDataDashboards {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) SetUrl(v string) *ListEnvironmentDashboardsResponseBodyDataDashboards {
	s.Url = &v
	return s
}

func (s *ListEnvironmentDashboardsResponseBodyDataDashboards) Validate() error {
	return dara.Validate(s)
}
