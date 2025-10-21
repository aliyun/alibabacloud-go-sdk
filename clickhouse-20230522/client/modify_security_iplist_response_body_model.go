// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifySecurityIPListResponseBodyData) *ModifySecurityIPListResponseBody
	GetData() *ModifySecurityIPListResponseBodyData
	SetRequestId(v string) *ModifySecurityIPListResponseBody
	GetRequestId() *string
}

type ModifySecurityIPListResponseBody struct {
	// The returned result.
	Data *ModifySecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponseBody) GetData() *ModifySecurityIPListResponseBodyData {
	return s.Data
}

func (s *ModifySecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIPListResponseBody) SetData(v *ModifySecurityIPListResponseBodyData) *ModifySecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetRequestId(v string) *ModifySecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySecurityIPListResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// 192.168.0.0/24,172.16.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The IP address type.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifySecurityIPListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *ModifySecurityIPListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySecurityIPListResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySecurityIPListResponseBodyData) GetGroupTag() *string {
	return s.GroupTag
}

func (s *ModifySecurityIPListResponseBodyData) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySecurityIPListResponseBodyData) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *ModifySecurityIPListResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifySecurityIPListResponseBodyData) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *ModifySecurityIPListResponseBodyData) SetDBInstanceID(v int32) *ModifySecurityIPListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetDBInstanceName(v string) *ModifySecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetGroupName(v string) *ModifySecurityIPListResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetGroupTag(v string) *ModifySecurityIPListResponseBodyData {
	s.GroupTag = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetSecurityIPList(v string) *ModifySecurityIPListResponseBodyData {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetSecurityIPType(v string) *ModifySecurityIPListResponseBodyData {
	s.SecurityIPType = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetTaskId(v int32) *ModifySecurityIPListResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) SetWhitelistNetType(v string) *ModifySecurityIPListResponseBodyData {
	s.WhitelistNetType = &v
	return s
}

func (s *ModifySecurityIPListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
