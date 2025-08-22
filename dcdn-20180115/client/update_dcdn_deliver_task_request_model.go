// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnDeliverTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliver(v string) *UpdateDcdnDeliverTaskRequest
	GetDeliver() *string
	SetDeliverId(v int64) *UpdateDcdnDeliverTaskRequest
	GetDeliverId() *int64
	SetDomainName(v string) *UpdateDcdnDeliverTaskRequest
	GetDomainName() *string
	SetName(v string) *UpdateDcdnDeliverTaskRequest
	GetName() *string
	SetReports(v string) *UpdateDcdnDeliverTaskRequest
	GetReports() *string
	SetSchedule(v string) *UpdateDcdnDeliverTaskRequest
	GetSchedule() *string
}

type UpdateDcdnDeliverTaskRequest struct {
	// The method that is used to send operations reports. Operations reports are sent to you only by email. The settings need to be escaped in JSON.
	//
	// example:
	//
	// {"email":{"subject":"the email subject","to":["username@example.com","username@example.com"]}}
	Deliver *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	// The ID of the tracking task that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	// The domain names from which the tracking task collects data. Separate domain names with commas (,). If you do not specify a domain name, the task collects data from all domain names that belong to your Alibaba Cloud account.
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
	// The operations reports that are tracked by the task. The data needs to be escaped in JSON.
	//
	// example:
	//
	// [{\\\\"reportId\\\\":2,\\\\"conditions\\\\":[{\\\\"field\\\\":\\\\"prov\\\\",\\\\"op\\\\":\\\\"in\\\\",\\\\"value\\\\":[\\\\"Heilongjiang\\\\",\\\\"Beijing\\\\"]}]}]
	Reports *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	// The parameters that specify the time interval at which the tracking task sends operations reports. The settings need to be escaped in JSON.
	//
	// example:
	//
	// {"schedName":"the name of the tracking task","description":"the description","crontab":"000\\*\\*?","frequency":"d","status":"enable","effectiveFrom":"2020-09-17T00:00:00Z","effectiveEnd":"2020-11-17T00:00:00Z"}
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateDcdnDeliverTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDeliverTaskRequest) GetDeliver() *string {
	return s.Deliver
}

func (s *UpdateDcdnDeliverTaskRequest) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *UpdateDcdnDeliverTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDcdnDeliverTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDcdnDeliverTaskRequest) GetReports() *string {
	return s.Reports
}

func (s *UpdateDcdnDeliverTaskRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateDcdnDeliverTaskRequest) SetDeliver(v string) *UpdateDcdnDeliverTaskRequest {
	s.Deliver = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetDeliverId(v int64) *UpdateDcdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetDomainName(v string) *UpdateDcdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetName(v string) *UpdateDcdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetReports(v string) *UpdateDcdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetSchedule(v string) *UpdateDcdnDeliverTaskRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) Validate() error {
	return dara.Validate(s)
}
