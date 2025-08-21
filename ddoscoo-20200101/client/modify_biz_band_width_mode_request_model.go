// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizBandWidthModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyBizBandWidthModeRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyBizBandWidthModeRequest
	GetMode() *string
}

type ModifyBizBandWidthModeRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-zvp2ay9b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the burstable clean bandwidth feature. Valid values:
	//
	// 	- **month**: the metering method of monthly 95th percentile
	//
	// 	- **day**: the metering method of daily 95th percentile
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyBizBandWidthModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizBandWidthModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBizBandWidthModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBizBandWidthModeRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyBizBandWidthModeRequest) SetInstanceId(v string) *ModifyBizBandWidthModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBizBandWidthModeRequest) SetMode(v string) *ModifyBizBandWidthModeRequest {
	s.Mode = &v
	return s
}

func (s *ModifyBizBandWidthModeRequest) Validate() error {
	return dara.Validate(s)
}
