// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPushAllTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushTaskRsp(v *ModifyPushAllTaskResponseBodyPushTaskRsp) *ModifyPushAllTaskResponseBody
	GetPushTaskRsp() *ModifyPushAllTaskResponseBodyPushTaskRsp
	SetRequestId(v string) *ModifyPushAllTaskResponseBody
	GetRequestId() *string
}

type ModifyPushAllTaskResponseBody struct {
	// The results of security check tasks.
	PushTaskRsp *ModifyPushAllTaskResponseBodyPushTaskRsp `json:"PushTaskRsp,omitempty" xml:"PushTaskRsp,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 91EB4AC7-7FEF-4C72-BE49-4414E459AEC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPushAllTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPushAllTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponseBody) GetPushTaskRsp() *ModifyPushAllTaskResponseBodyPushTaskRsp {
	return s.PushTaskRsp
}

func (s *ModifyPushAllTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPushAllTaskResponseBody) SetPushTaskRsp(v *ModifyPushAllTaskResponseBodyPushTaskRsp) *ModifyPushAllTaskResponseBody {
	s.PushTaskRsp = v
	return s
}

func (s *ModifyPushAllTaskResponseBody) SetRequestId(v string) *ModifyPushAllTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPushAllTaskResponseBody) Validate() error {
	if s.PushTaskRsp != nil {
		if err := s.PushTaskRsp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPushAllTaskResponseBodyPushTaskRsp struct {
	// The information about the server on which security check tasks failed.
	PushTaskResultList []*ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList `json:"PushTaskResultList,omitempty" xml:"PushTaskResultList,omitempty" type:"Repeated"`
}

func (s ModifyPushAllTaskResponseBodyPushTaskRsp) String() string {
	return dara.Prettify(s)
}

func (s ModifyPushAllTaskResponseBodyPushTaskRsp) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRsp) GetPushTaskResultList() []*ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	return s.PushTaskResultList
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRsp) SetPushTaskResultList(v []*ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) *ModifyPushAllTaskResponseBodyPushTaskRsp {
	s.PushTaskResultList = v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRsp) Validate() error {
	if s.PushTaskResultList != nil {
		for _, item := range s.PushTaskResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList struct {
	// The ID of the server group to which the server belongs.
	//
	// example:
	//
	// 226
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-wz9f7wlklxqnvdk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// TestInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 127.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The message that describes the security check failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the Security Center agent is online. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  If the Security Center agent of the server is offline, Security Center does not protect the server.
	//
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The operating system version of the server.
	//
	// example:
	//
	// linux
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether the security check task is successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 5493fe42-61f5-4627-9aa2-8c449bbe****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetIp() *string {
	return s.Ip
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetMessage() *string {
	return s.Message
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetOnline() *bool {
	return s.Online
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetOsVersion() *string {
	return s.OsVersion
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetRegion() *string {
	return s.Region
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetGroupId(v int64) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.GroupId = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetInstanceId(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.InstanceId = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetInstanceName(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.InstanceName = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetIp(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Ip = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetMessage(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Message = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetOnline(v bool) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Online = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetOsVersion(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.OsVersion = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetRegion(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Region = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetSuccess(v bool) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Success = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetUuid(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Uuid = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) Validate() error {
	return dara.Validate(s)
}
