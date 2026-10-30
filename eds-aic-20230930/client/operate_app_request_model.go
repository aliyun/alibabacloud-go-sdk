// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *OperateAppRequest
	GetAppId() *int32
	SetInstanceIdList(v []*string) *OperateAppRequest
	GetInstanceIdList() []*string
	SetOperateType(v string) *OperateAppRequest
	GetOperateType() *string
}

type OperateAppRequest struct {
	// The application ID.
	//
	// example:
	//
	// 1234
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The operation type. Supported operations include starting, stopping, and restarting.
	//
	// example:
	//
	// start
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s OperateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAppRequest) GoString() string {
	return s.String()
}

func (s *OperateAppRequest) GetAppId() *int32 {
	return s.AppId
}

func (s *OperateAppRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *OperateAppRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateAppRequest) SetAppId(v int32) *OperateAppRequest {
	s.AppId = &v
	return s
}

func (s *OperateAppRequest) SetInstanceIdList(v []*string) *OperateAppRequest {
	s.InstanceIdList = v
	return s
}

func (s *OperateAppRequest) SetOperateType(v string) *OperateAppRequest {
	s.OperateType = &v
	return s
}

func (s *OperateAppRequest) Validate() error {
	return dara.Validate(s)
}
