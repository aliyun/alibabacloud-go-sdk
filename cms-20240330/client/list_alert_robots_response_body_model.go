// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRobotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAlertRobotsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertRobotsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAlertRobotsResponseBody
	GetRequestId() *string
	SetRobots(v []*ListAlertRobotsResponseBodyRobots) *ListAlertRobotsResponseBody
	GetRobots() []*ListAlertRobotsResponseBodyRobots
	SetTotal(v int64) *ListAlertRobotsResponseBody
	GetTotal() *int64
}

type ListAlertRobotsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Robots    []*ListAlertRobotsResponseBodyRobots `json:"robots,omitempty" xml:"robots,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertRobotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRobotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertRobotsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertRobotsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertRobotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertRobotsResponseBody) GetRobots() []*ListAlertRobotsResponseBodyRobots {
	return s.Robots
}

func (s *ListAlertRobotsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlertRobotsResponseBody) SetPageNumber(v int64) *ListAlertRobotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertRobotsResponseBody) SetPageSize(v int64) *ListAlertRobotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertRobotsResponseBody) SetRequestId(v string) *ListAlertRobotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertRobotsResponseBody) SetRobots(v []*ListAlertRobotsResponseBodyRobots) *ListAlertRobotsResponseBody {
	s.Robots = v
	return s
}

func (s *ListAlertRobotsResponseBody) SetTotal(v int64) *ListAlertRobotsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAlertRobotsResponseBody) Validate() error {
	if s.Robots != nil {
		for _, item := range s.Robots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertRobotsResponseBodyRobots struct {
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test
	RobotId *string `json:"robotId,omitempty" xml:"robotId,omitempty"`
	// example:
	//
	// DING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=*******
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAlertRobotsResponseBodyRobots) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRobotsResponseBodyRobots) GoString() string {
	return s.String()
}

func (s *ListAlertRobotsResponseBodyRobots) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *ListAlertRobotsResponseBodyRobots) GetLang() *string {
	return s.Lang
}

func (s *ListAlertRobotsResponseBodyRobots) GetName() *string {
	return s.Name
}

func (s *ListAlertRobotsResponseBodyRobots) GetRobotId() *string {
	return s.RobotId
}

func (s *ListAlertRobotsResponseBodyRobots) GetType() *string {
	return s.Type
}

func (s *ListAlertRobotsResponseBodyRobots) GetUrl() *string {
	return s.Url
}

func (s *ListAlertRobotsResponseBodyRobots) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertRobotsResponseBodyRobots) SetDigitalEmployeeName(v string) *ListAlertRobotsResponseBodyRobots {
	s.DigitalEmployeeName = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) SetLang(v string) *ListAlertRobotsResponseBodyRobots {
	s.Lang = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) SetName(v string) *ListAlertRobotsResponseBodyRobots {
	s.Name = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) SetRobotId(v string) *ListAlertRobotsResponseBodyRobots {
	s.RobotId = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) SetType(v string) *ListAlertRobotsResponseBodyRobots {
	s.Type = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) SetUrl(v string) *ListAlertRobotsResponseBodyRobots {
	s.Url = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) SetWorkspace(v string) *ListAlertRobotsResponseBodyRobots {
	s.Workspace = &v
	return s
}

func (s *ListAlertRobotsResponseBodyRobots) Validate() error {
	return dara.Validate(s)
}
