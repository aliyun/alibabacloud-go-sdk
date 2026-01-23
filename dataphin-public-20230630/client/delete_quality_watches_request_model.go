// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityWatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteQualityWatchesRequestDeleteCommand) *DeleteQualityWatchesRequest
	GetDeleteCommand() *DeleteQualityWatchesRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteQualityWatchesRequest
	GetOpTenantId() *int64
}

type DeleteQualityWatchesRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteQualityWatchesRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteQualityWatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityWatchesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityWatchesRequest) GetDeleteCommand() *DeleteQualityWatchesRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteQualityWatchesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualityWatchesRequest) SetDeleteCommand(v *DeleteQualityWatchesRequestDeleteCommand) *DeleteQualityWatchesRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteQualityWatchesRequest) SetOpTenantId(v int64) *DeleteQualityWatchesRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityWatchesRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteQualityWatchesRequestDeleteCommand struct {
	// This parameter is required.
	WatchIdList []*int64 `json:"WatchIdList,omitempty" xml:"WatchIdList,omitempty" type:"Repeated"`
}

func (s DeleteQualityWatchesRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityWatchesRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteQualityWatchesRequestDeleteCommand) GetWatchIdList() []*int64 {
	return s.WatchIdList
}

func (s *DeleteQualityWatchesRequestDeleteCommand) SetWatchIdList(v []*int64) *DeleteQualityWatchesRequestDeleteCommand {
	s.WatchIdList = v
	return s
}

func (s *DeleteQualityWatchesRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
