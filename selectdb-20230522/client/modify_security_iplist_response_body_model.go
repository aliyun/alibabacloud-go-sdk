// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifySecurityIPListResponseBody
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifySecurityIPListResponseBody
	GetGroupName() *string
	SetGroupTag(v string) *ModifySecurityIPListResponseBody
	GetGroupTag() *string
	SetRequestId(v string) *ModifySecurityIPListResponseBody
	GetRequestId() *string
	SetSecurityIPList(v string) *ModifySecurityIPListResponseBody
	GetSecurityIPList() *string
	SetSecurityIPType(v string) *ModifySecurityIPListResponseBody
	GetSecurityIPType() *string
	SetTaskId(v int64) *ModifySecurityIPListResponseBody
	GetTaskId() *int64
	SetWhitelistNetType(v string) *ModifySecurityIPListResponseBody
	GetWhitelistNetType() *string
}

type ModifySecurityIPListResponseBody struct {
	// The name of the instance.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist.
	//
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist.
	//
	// example:
	//
	// grouptag
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 195F64C2-8F11-532B-A436-FC08A221D756
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP addresses in the whitelist of the instance. Multiple IP addresses are separated by commas (,).
	//
	// example:
	//
	// 127.0.XX.XX,127.2.XX.XX
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
	// 479095561
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifySecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySecurityIPListResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySecurityIPListResponseBody) GetGroupTag() *string {
	return s.GroupTag
}

func (s *ModifySecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIPListResponseBody) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySecurityIPListResponseBody) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *ModifySecurityIPListResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifySecurityIPListResponseBody) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *ModifySecurityIPListResponseBody) SetDBInstanceName(v string) *ModifySecurityIPListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetGroupName(v string) *ModifySecurityIPListResponseBody {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetGroupTag(v string) *ModifySecurityIPListResponseBody {
	s.GroupTag = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetRequestId(v string) *ModifySecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetSecurityIPList(v string) *ModifySecurityIPListResponseBody {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetSecurityIPType(v string) *ModifySecurityIPListResponseBody {
	s.SecurityIPType = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetTaskId(v int64) *ModifySecurityIPListResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) SetWhitelistNetType(v string) *ModifySecurityIPListResponseBody {
	s.WhitelistNetType = &v
	return s
}

func (s *ModifySecurityIPListResponseBody) Validate() error {
	return dara.Validate(s)
}
