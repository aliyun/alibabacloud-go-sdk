// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQpsModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyQpsModeRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyQpsModeRequest
	GetMode() *string
}

type ModifyQpsModeRequest struct {
	// The region ID of the Anti-DDoS Pro instance.
	//
	// >  You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-7e225i41****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of QPS. Valid values:
	//
	// 	- **month**: monthly 95th percentile QPS.
	//
	// 	- **day**: daily 95th percentile QPS.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyQpsModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQpsModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyQpsModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyQpsModeRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyQpsModeRequest) SetInstanceId(v string) *ModifyQpsModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyQpsModeRequest) SetMode(v string) *ModifyQpsModeRequest {
	s.Mode = &v
	return s
}

func (s *ModifyQpsModeRequest) Validate() error {
	return dara.Validate(s)
}
