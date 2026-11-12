// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DeleteCommandRequest
	GetCommandId() *string
	SetRegionId(v string) *DeleteCommandRequest
	GetRegionId() *string
}

type DeleteCommandRequest struct {
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the available regions for Simple Application Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DeleteCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCommandRequest) SetCommandId(v string) *DeleteCommandRequest {
	s.CommandId = &v
	return s
}

func (s *DeleteCommandRequest) SetRegionId(v string) *DeleteCommandRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCommandRequest) Validate() error {
	return dara.Validate(s)
}
