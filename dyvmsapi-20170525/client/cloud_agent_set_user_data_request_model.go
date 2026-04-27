// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentSetUserDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudAgentSetUserDataRequest
	GetCno() *string
	SetDirection(v int64) *CloudAgentSetUserDataRequest
	GetDirection() *int64
	SetEnterpriseId(v int64) *CloudAgentSetUserDataRequest
	GetEnterpriseId() *int64
	SetUserData(v string) *CloudAgentSetUserDataRequest
	GetUserData() *string
}

type CloudAgentSetUserDataRequest struct {
	// 座席工号；取值3-10位正整数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 随路数据方向；取值说明：1： 座席侧，2： 非座席侧
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Direction *int64 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// json格式字符串,传入的值会打入通道变量,格式：json字符串{"key":"value"} 随路数据数据格式：key=value,不支持数组，嵌套。
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key":"value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CloudAgentSetUserDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentSetUserDataRequest) GoString() string {
	return s.String()
}

func (s *CloudAgentSetUserDataRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudAgentSetUserDataRequest) GetDirection() *int64 {
	return s.Direction
}

func (s *CloudAgentSetUserDataRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAgentSetUserDataRequest) GetUserData() *string {
	return s.UserData
}

func (s *CloudAgentSetUserDataRequest) SetCno(v string) *CloudAgentSetUserDataRequest {
	s.Cno = &v
	return s
}

func (s *CloudAgentSetUserDataRequest) SetDirection(v int64) *CloudAgentSetUserDataRequest {
	s.Direction = &v
	return s
}

func (s *CloudAgentSetUserDataRequest) SetEnterpriseId(v int64) *CloudAgentSetUserDataRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentSetUserDataRequest) SetUserData(v string) *CloudAgentSetUserDataRequest {
	s.UserData = &v
	return s
}

func (s *CloudAgentSetUserDataRequest) Validate() error {
	return dara.Validate(s)
}
