// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPrepayInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConvertPrepayInstanceRequest
	GetInstanceId() *string
	SetRegion(v string) *ConvertPrepayInstanceRequest
	GetRegion() *string
}

type ConvertPrepayInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertPrepayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertPrepayInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertPrepayInstanceRequest) SetInstanceId(v string) *ConvertPrepayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertPrepayInstanceRequest) SetRegion(v string) *ConvertPrepayInstanceRequest {
	s.Region = &v
	return s
}

func (s *ConvertPrepayInstanceRequest) Validate() error {
	return dara.Validate(s)
}
