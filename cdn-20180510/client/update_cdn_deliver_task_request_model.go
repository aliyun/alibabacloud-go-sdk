// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCdnDeliverTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliver(v string) *UpdateCdnDeliverTaskRequest
	GetDeliver() *string
	SetDeliverId(v int64) *UpdateCdnDeliverTaskRequest
	GetDeliverId() *int64
	SetDomainName(v string) *UpdateCdnDeliverTaskRequest
	GetDomainName() *string
	SetName(v string) *UpdateCdnDeliverTaskRequest
	GetName() *string
	SetReports(v string) *UpdateCdnDeliverTaskRequest
	GetReports() *string
	SetSchedule(v string) *UpdateCdnDeliverTaskRequest
	GetSchedule() *string
}

type UpdateCdnDeliverTaskRequest struct {
	// The method that is used to send operations reports. Operations reports are sent to you only by email. The settings must be escaped in JSON.
	//
	// example:
	//
	// {\\\\"email\\\\":{\\\\"subject\\\\":\\\\"The email subject\\\\",\\\\"to\\\\":[\\\\"songmingyuan@alibaba-inc.com\\\\",\\\\"songmingyuan@alibaba-inc.com\\\\"]}}"
	Deliver *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	// The ID of the tracking task that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	// The domain name that you want to track. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,). If you do not specify a domain name, the task collects data from all domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the tracking task.
	//
	// example:
	//
	// Domain name report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations reports that are tracked by the task. The data must be escaped in JSON.
	//
	// example:
	//
	// [{\\\\"reportId\\\\":1,\\\\"conditions\\\\":[{\\\\"field\\\\":\\\\"prov\\\\",\\\\"op\\\\":\\\\"in\\\\",\\\\"value\\\\":[\\\\"Heilongjiang\\\\",\\\\"Beijing\\\\"]}]}]
	Reports *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	// The parameters that specify the time interval at which the tracking task sends operations reports. The settings must be escaped in JSON.
	//
	// example:
	//
	// "{\\\\"schedName\\\\":\\\\"The name of the tracking task\\\\",\\\\"description\\\\":\\\\"The description\\\\",\\\\"crontab\\\\":\\\\"000\\*\\*?\\\\",\\\\"frequency\\\\":\\\\"d\\\\",\\\\"status\\\\":\\\\"enable\\\\",\\\\"effectiveFrom\\\\":\\\\"2020-09-17T00:00:00Z\\\\",\\\\"effectiveEnd\\\\":\\\\"2020-11-17T00:00:00Z\\\\"}"
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateCdnDeliverTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskRequest) GetDeliver() *string {
	return s.Deliver
}

func (s *UpdateCdnDeliverTaskRequest) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *UpdateCdnDeliverTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateCdnDeliverTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCdnDeliverTaskRequest) GetReports() *string {
	return s.Reports
}

func (s *UpdateCdnDeliverTaskRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateCdnDeliverTaskRequest) SetDeliver(v string) *UpdateCdnDeliverTaskRequest {
	s.Deliver = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetDeliverId(v int64) *UpdateCdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetDomainName(v string) *UpdateCdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetName(v string) *UpdateCdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetReports(v string) *UpdateCdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetSchedule(v string) *UpdateCdnDeliverTaskRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) Validate() error {
	return dara.Validate(s)
}
