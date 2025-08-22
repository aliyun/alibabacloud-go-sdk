// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnDeliverTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliver(v string) *CreateDcdnDeliverTaskRequest
	GetDeliver() *string
	SetDomainName(v string) *CreateDcdnDeliverTaskRequest
	GetDomainName() *string
	SetName(v string) *CreateDcdnDeliverTaskRequest
	GetName() *string
	SetReports(v string) *CreateDcdnDeliverTaskRequest
	GetReports() *string
	SetSchedule(v string) *CreateDcdnDeliverTaskRequest
	GetSchedule() *string
}

type CreateDcdnDeliverTaskRequest struct {
	// The method that is used to send operations reports. Operations reports are sent to you only by email. The settings must be escaped in JSON.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"email":{"subject":"the email subject","to":["username@example.com","username@example.org"]}}
	Deliver *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	// The domain names to be tracked. Separate multiple domain names with commas (,). You can specify up to 500 domain names. If you want to specify more than 500 domain names, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
	//
	// > If you do not specify a domain name, the tracking task is created for all domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the tracking task.
	//
	// This parameter is required.
	//
	// example:
	//
	// Domain name report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations reports that are tracked by the task. The data must be escaped in JSON.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\\\"reportId\\\\":2,\\\\"conditions\\\\":[{\\\\"field\\\\":\\\\"prov\\\\",\\\\"op\\\\":\\\\"in\\\\",\\\\"value\\\\":[\\\\"Heilongjiang\\\\",\\\\"Beijing\\\\"]}]}]
	Reports *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	// The parameters that specify the time interval at which the tracking task sends operations reports. The settings must be escaped in JSON.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"schedName":"the name of the tracking task","description":"the description","crontab":"000\\*\\*?","frequency":"d","status":"enable","effectiveFrom":"2020-09-17T00:00:00Z","effectiveEnd":"2020-11-17T00:00:00Z"}
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateDcdnDeliverTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnDeliverTaskRequest) GetDeliver() *string {
	return s.Deliver
}

func (s *CreateDcdnDeliverTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateDcdnDeliverTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateDcdnDeliverTaskRequest) GetReports() *string {
	return s.Reports
}

func (s *CreateDcdnDeliverTaskRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateDcdnDeliverTaskRequest) SetDeliver(v string) *CreateDcdnDeliverTaskRequest {
	s.Deliver = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetDomainName(v string) *CreateDcdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetName(v string) *CreateDcdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetReports(v string) *CreateDcdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetSchedule(v string) *CreateDcdnDeliverTaskRequest {
	s.Schedule = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) Validate() error {
	return dara.Validate(s)
}
