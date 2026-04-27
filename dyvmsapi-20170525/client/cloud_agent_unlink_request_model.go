// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentUnlinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *CloudAgentUnlinkRequest
	GetCno() *string
	SetEnterpriseId(v int64) *CloudAgentUnlinkRequest
	GetEnterpriseId() *int64
	SetRequestUniqueId(v string) *CloudAgentUnlinkRequest
	GetRequestUniqueId() *string
	SetSide(v int64) *CloudAgentUnlinkRequest
	GetSide() *int64
	SetUniqueId(v string) *CloudAgentUnlinkRequest
	GetUniqueId() *string
}

type CloudAgentUnlinkRequest struct {
	// 座席工号；取值3-10位正整数，参数cno,uniqueId和requestUniqueId三选一
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 请求唯一标识；参数cno,uniqueId和requestUniqueId三选一。注意：requestUniqueId挂机功能只在呼叫类型为webcall时生效，且需提前在管理后台将功能开启
	//
	// example:
	//
	// 1233456789
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// 是否座席侧挂机；取值范围：1,2（数字含义，座席侧，非座席侧)，默认座席侧
	//
	// example:
	//
	// 1
	Side *int64 `json:"Side,omitempty" xml:"Side,omitempty"`
	// 通话唯一标识；参数cno,uniqueId和requestUniqueId三选一
	//
	// example:
	//
	// 1233456789
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s CloudAgentUnlinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentUnlinkRequest) GoString() string {
	return s.String()
}

func (s *CloudAgentUnlinkRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudAgentUnlinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAgentUnlinkRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudAgentUnlinkRequest) GetSide() *int64 {
	return s.Side
}

func (s *CloudAgentUnlinkRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudAgentUnlinkRequest) SetCno(v string) *CloudAgentUnlinkRequest {
	s.Cno = &v
	return s
}

func (s *CloudAgentUnlinkRequest) SetEnterpriseId(v int64) *CloudAgentUnlinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentUnlinkRequest) SetRequestUniqueId(v string) *CloudAgentUnlinkRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudAgentUnlinkRequest) SetSide(v int64) *CloudAgentUnlinkRequest {
	s.Side = &v
	return s
}

func (s *CloudAgentUnlinkRequest) SetUniqueId(v string) *CloudAgentUnlinkRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudAgentUnlinkRequest) Validate() error {
	return dara.Validate(s)
}
