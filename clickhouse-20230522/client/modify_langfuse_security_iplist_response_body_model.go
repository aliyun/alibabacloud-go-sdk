// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseSecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyLangfuseSecurityIPListResponseBodyData) *ModifyLangfuseSecurityIPListResponseBody
	GetData() *ModifyLangfuseSecurityIPListResponseBodyData
	SetRequestId(v string) *ModifyLangfuseSecurityIPListResponseBody
	GetRequestId() *string
}

type ModifyLangfuseSecurityIPListResponseBody struct {
	// The returned result.
	Data *ModifyLangfuseSecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLangfuseSecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseSecurityIPListResponseBody) GetData() *ModifyLangfuseSecurityIPListResponseBodyData {
	return s.Data
}

func (s *ModifyLangfuseSecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLangfuseSecurityIPListResponseBody) SetData(v *ModifyLangfuseSecurityIPListResponseBodyData) *ModifyLangfuseSecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBody) SetRequestId(v string) *ModifyLangfuseSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyLangfuseSecurityIPListResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// 12345
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The Langfuse instance ID.
	//
	// example:
	//
	// lfs-*****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The group tag.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP whitelist.
	//
	// example:
	//
	// 192.168.0.0/24,172.16.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The IP address type. The value is fixed to IPv4. IPv6 is not supported.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 100001080
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifyLangfuseSecurityIPListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseSecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetGroupTag() *string {
	return s.GroupTag
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetDBInstanceID(v int32) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetDBInstanceName(v string) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetGroupName(v string) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetGroupTag(v string) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.GroupTag = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetSecurityIPList(v string) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetSecurityIPType(v string) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.SecurityIPType = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetTaskId(v int32) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) SetWhitelistNetType(v string) *ModifyLangfuseSecurityIPListResponseBodyData {
	s.WhitelistNetType = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
