// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIMRobotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *DescribeIMRobotsResponseBodyPageBean) *DescribeIMRobotsResponseBody
	GetPageBean() *DescribeIMRobotsResponseBodyPageBean
	SetRequestId(v string) *DescribeIMRobotsResponseBody
	GetRequestId() *string
}

type DescribeIMRobotsResponseBody struct {
	// The returned objects.
	PageBean *DescribeIMRobotsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4D6C358A-A58B-4F4B-94CE-F5AAF023****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIMRobotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMRobotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBody) GetPageBean() *DescribeIMRobotsResponseBodyPageBean {
	return s.PageBean
}

func (s *DescribeIMRobotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIMRobotsResponseBody) SetPageBean(v *DescribeIMRobotsResponseBodyPageBean) *DescribeIMRobotsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeIMRobotsResponseBody) SetRequestId(v string) *DescribeIMRobotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIMRobotsResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIMRobotsResponseBodyPageBean struct {
	// The queried IM chatbots.
	AlertIMRobots []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobots `json:"AlertIMRobots,omitempty" xml:"AlertIMRobots,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of IM chatbots returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of queried IM chatbots.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeIMRobotsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMRobotsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBodyPageBean) GetAlertIMRobots() []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	return s.AlertIMRobots
}

func (s *DescribeIMRobotsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *DescribeIMRobotsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *DescribeIMRobotsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetAlertIMRobots(v []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) *DescribeIMRobotsResponseBodyPageBean {
	s.AlertIMRobots = v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetPage(v int64) *DescribeIMRobotsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetSize(v int64) *DescribeIMRobotsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetTotal(v int64) *DescribeIMRobotsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) Validate() error {
	if s.AlertIMRobots != nil {
		for _, item := range s.AlertIMRobots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIMRobotsResponseBodyPageBeanAlertIMRobots struct {
	// The time when the IM chatbot was created.
	//
	// example:
	//
	// 2023-01-16 17:21:48
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether daily statistics are sent. Valid values:
	//
	// 	- `false` (default): Daily statistics are not sent.
	//
	// 	- `true`: Daily statistics are sent.
	//
	// example:
	//
	// true
	DailyNoc *bool `json:"DailyNoc,omitempty" xml:"DailyNoc,omitempty"`
	// The point in time at which the daily statistics are sent. The information that ARMS sends at the specified points in time includes the total number of alerts generated on the current day, the number of cleared alerts, and the number of alerts to be cleared.
	//
	// example:
	//
	// 09:30,17:30
	DailyNocTime *string `json:"DailyNocTime,omitempty" xml:"DailyNocTime,omitempty"`
	// The signature key of DingTalk. If you specify a signature key, DingTalk authentication is performed by using the signature key. If you do not specify a signature key, a whitelist is used for authentication by default. The keyword of the whitelist is **Alert**.
	//
	// example:
	//
	// ******
	DingSignKey *string `json:"DingSignKey,omitempty" xml:"DingSignKey,omitempty"`
	// The notification policies.
	DispatchRules []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	// The webhook URL of the IM chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121******
	RobotAddr *string `json:"RobotAddr,omitempty" xml:"RobotAddr,omitempty"`
	// The ID of the IM chatbot.
	//
	// example:
	//
	// 123
	RobotId *float32 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// The name of the IM chatbot.
	//
	// example:
	//
	// Chatbot name
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// The type of the IM chatbot. Valid values:
	//
	// 	- `dingding`: DingTalk chatbot
	//
	// 	- `wechat`: WeCom chatbot
	//
	// example:
	//
	// dingding
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetDailyNoc() *bool {
	return s.DailyNoc
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetDailyNocTime() *string {
	return s.DailyNocTime
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetDingSignKey() *string {
	return s.DingSignKey
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetDispatchRules() []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules {
	return s.DispatchRules
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetRobotAddr() *string {
	return s.RobotAddr
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetRobotId() *float32 {
	return s.RobotId
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetRobotName() *string {
	return s.RobotName
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GetType() *string {
	return s.Type
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetCreateTime(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.CreateTime = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetDailyNoc(v bool) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.DailyNoc = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetDailyNocTime(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.DailyNocTime = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetDingSignKey(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.DingSignKey = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetDispatchRules(v []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.DispatchRules = v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetRobotAddr(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.RobotAddr = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetRobotId(v float32) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.RobotId = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetRobotName(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.RobotName = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetType(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.Type = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) Validate() error {
	if s.DispatchRules != nil {
		for _, item := range s.DispatchRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules struct {
	// The ID of the notification policy.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// Notification policy test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) GetId() *int64 {
	return s.Id
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) GetName() *string {
	return s.Name
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) SetId(v int64) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules {
	s.Id = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) SetName(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules {
	s.Name = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobotsDispatchRules) Validate() error {
	return dara.Validate(s)
}
