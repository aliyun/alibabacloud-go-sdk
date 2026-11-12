// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplateApplyResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeFirewallTemplateApplyResultsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeFirewallTemplateApplyResultsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeFirewallTemplateApplyResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeFirewallTemplateApplyResultsResponseBody
	GetTotalCount() *string
	SetData(v []*DescribeFirewallTemplateApplyResultsResponseBodyData) *DescribeFirewallTemplateApplyResultsResponseBody
	GetData() []*DescribeFirewallTemplateApplyResultsResponseBodyData
}

type DescribeFirewallTemplateApplyResultsResponseBody struct {
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of firewall rules.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The returned results.
	Data []*DescribeFirewallTemplateApplyResultsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeFirewallTemplateApplyResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) GetData() []*DescribeFirewallTemplateApplyResultsResponseBodyData {
	return s.Data
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetPageNumber(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetPageSize(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetRequestId(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetTotalCount(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetData(v []*DescribeFirewallTemplateApplyResultsResponseBodyData) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) Validate() error {
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

type DescribeFirewallTemplateApplyResultsResponseBodyData struct {
	// The time when the firewall template application task was executed, in the format of EEE MMM dd HH:mm:ss zzz yyyy.
	//
	// example:
	//
	// Tue May 14 11:53:07 CST 2024
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of instances to which the firewall template failed to be applied.
	//
	// example:
	//
	// 0
	FailedCount *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The firewall template ID.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The results of applying the firewall template to instances.
	InstanceApplyResults []*DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults `json:"InstanceApplyResults,omitempty" xml:"InstanceApplyResults,omitempty" type:"Repeated"`
	// The status of applying the firewall template to all instances. Valid values:
	//
	// - Running: The firewall template is being applied.
	//
	// - Failed: The firewall template failed to be applied to all instances.
	//
	// - Success: The firewall template was applied.
	//
	// - PartFailed: The firewall template was partially applied. Some instances failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template application task.
	//
	// example:
	//
	// aft-ikgt0bynitvs3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of instances to which the firewall template is applied.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetFailedCount() *string {
	return s.FailedCount
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetInstanceApplyResults() []*DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults {
	return s.InstanceApplyResults
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetCreateTime(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetFailedCount(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetFirewallTemplateId(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetInstanceApplyResults(v []*DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.InstanceApplyResults = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetStatus(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetTaskId(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetTotalCount(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) Validate() error {
	if s.InstanceApplyResults != nil {
		for _, item := range s.InstanceApplyResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults struct {
	// The simple application server instance ID.
	//
	// example:
	//
	// 33774babccc84003a2b1ad47e8001233
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of applying the firewall template to a specific instance. Valid values:
	//
	// - Running: The firewall template is being applied.
	//
	// - Failed: The firewall template failed to be applied.
	//
	// - Success: The firewall template was applied.
	//
	// - PartFailed: The firewall template was partially applied. Some rules failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) GetStatus() *string {
	return s.Status
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) SetInstanceId(v string) *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults {
	s.InstanceId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) SetStatus(v string) *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults {
	s.Status = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) Validate() error {
	return dara.Validate(s)
}
