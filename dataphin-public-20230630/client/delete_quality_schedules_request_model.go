// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualitySchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteQualitySchedulesRequestDeleteCommand) *DeleteQualitySchedulesRequest
	GetDeleteCommand() *DeleteQualitySchedulesRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteQualitySchedulesRequest
	GetOpTenantId() *int64
}

type DeleteQualitySchedulesRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteQualitySchedulesRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteQualitySchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualitySchedulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualitySchedulesRequest) GetDeleteCommand() *DeleteQualitySchedulesRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteQualitySchedulesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualitySchedulesRequest) SetDeleteCommand(v *DeleteQualitySchedulesRequestDeleteCommand) *DeleteQualitySchedulesRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteQualitySchedulesRequest) SetOpTenantId(v int64) *DeleteQualitySchedulesRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualitySchedulesRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteQualitySchedulesRequestDeleteCommand struct {
	// This parameter is required.
	ScheduleIdList []*int64 `json:"ScheduleIdList,omitempty" xml:"ScheduleIdList,omitempty" type:"Repeated"`
}

func (s DeleteQualitySchedulesRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualitySchedulesRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteQualitySchedulesRequestDeleteCommand) GetScheduleIdList() []*int64 {
	return s.ScheduleIdList
}

func (s *DeleteQualitySchedulesRequestDeleteCommand) SetScheduleIdList(v []*int64) *DeleteQualitySchedulesRequestDeleteCommand {
	s.ScheduleIdList = v
	return s
}

func (s *DeleteQualitySchedulesRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
