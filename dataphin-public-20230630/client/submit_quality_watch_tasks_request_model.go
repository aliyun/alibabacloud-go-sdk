// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityWatchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitQualityWatchTasksRequest
	GetOpTenantId() *int64
	SetSubmitCommand(v *SubmitQualityWatchTasksRequestSubmitCommand) *SubmitQualityWatchTasksRequest
	GetSubmitCommand() *SubmitQualityWatchTasksRequestSubmitCommand
}

type SubmitQualityWatchTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitCommand *SubmitQualityWatchTasksRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitQualityWatchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityWatchTasksRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityWatchTasksRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitQualityWatchTasksRequest) GetSubmitCommand() *SubmitQualityWatchTasksRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitQualityWatchTasksRequest) SetOpTenantId(v int64) *SubmitQualityWatchTasksRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitQualityWatchTasksRequest) SetSubmitCommand(v *SubmitQualityWatchTasksRequestSubmitCommand) *SubmitQualityWatchTasksRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitQualityWatchTasksRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitQualityWatchTasksRequestSubmitCommand struct {
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// ds=${yyyyMMdd}
	PartitionExpression *string `json:"PartitionExpression,omitempty" xml:"PartitionExpression,omitempty"`
	// This parameter is required.
	WatchIdList []*int64 `json:"WatchIdList,omitempty" xml:"WatchIdList,omitempty" type:"Repeated"`
}

func (s SubmitQualityWatchTasksRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityWatchTasksRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) GetBizDate() *string {
	return s.BizDate
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) GetPartitionExpression() *string {
	return s.PartitionExpression
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) GetWatchIdList() []*int64 {
	return s.WatchIdList
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) SetBizDate(v string) *SubmitQualityWatchTasksRequestSubmitCommand {
	s.BizDate = &v
	return s
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) SetPartitionExpression(v string) *SubmitQualityWatchTasksRequestSubmitCommand {
	s.PartitionExpression = &v
	return s
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) SetWatchIdList(v []*int64) *SubmitQualityWatchTasksRequestSubmitCommand {
	s.WatchIdList = v
	return s
}

func (s *SubmitQualityWatchTasksRequestSubmitCommand) Validate() error {
	return dara.Validate(s)
}
