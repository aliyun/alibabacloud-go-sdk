// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteBizMetricCommand(v *DeleteBizMetricRequestDeleteBizMetricCommand) *DeleteBizMetricRequest
	GetDeleteBizMetricCommand() *DeleteBizMetricRequestDeleteBizMetricCommand
	SetOpTenantId(v int64) *DeleteBizMetricRequest
	GetOpTenantId() *int64
}

type DeleteBizMetricRequest struct {
	// This parameter is required.
	DeleteBizMetricCommand *DeleteBizMetricRequestDeleteBizMetricCommand `json:"DeleteBizMetricCommand,omitempty" xml:"DeleteBizMetricCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteBizMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizMetricRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizMetricRequest) GetDeleteBizMetricCommand() *DeleteBizMetricRequestDeleteBizMetricCommand {
	return s.DeleteBizMetricCommand
}

func (s *DeleteBizMetricRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteBizMetricRequest) SetDeleteBizMetricCommand(v *DeleteBizMetricRequestDeleteBizMetricCommand) *DeleteBizMetricRequest {
	s.DeleteBizMetricCommand = v
	return s
}

func (s *DeleteBizMetricRequest) SetOpTenantId(v int64) *DeleteBizMetricRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBizMetricRequest) Validate() error {
	if s.DeleteBizMetricCommand != nil {
		if err := s.DeleteBizMetricCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteBizMetricRequestDeleteBizMetricCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// Metric1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteBizMetricRequestDeleteBizMetricCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizMetricRequestDeleteBizMetricCommand) GoString() string {
	return s.String()
}

func (s *DeleteBizMetricRequestDeleteBizMetricCommand) GetName() *string {
	return s.Name
}

func (s *DeleteBizMetricRequestDeleteBizMetricCommand) SetName(v string) *DeleteBizMetricRequestDeleteBizMetricCommand {
	s.Name = &v
	return s
}

func (s *DeleteBizMetricRequestDeleteBizMetricCommand) Validate() error {
	return dara.Validate(s)
}
