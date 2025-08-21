// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageCount(v int64) *DescribeParentPlatformsResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeParentPlatformsResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeParentPlatformsResponseBody
	GetPageSize() *int64
	SetPlatforms(v []*DescribeParentPlatformsResponseBodyPlatforms) *DescribeParentPlatformsResponseBody
	GetPlatforms() []*DescribeParentPlatformsResponseBodyPlatforms
	SetRequestId(v string) *DescribeParentPlatformsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeParentPlatformsResponseBody
	GetTotalCount() *int64
}

type DescribeParentPlatformsResponseBody struct {
	// example:
	//
	// 1
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize  *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platforms []*DescribeParentPlatformsResponseBodyPlatforms `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParentPlatformsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeParentPlatformsResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeParentPlatformsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeParentPlatformsResponseBody) GetPlatforms() []*DescribeParentPlatformsResponseBodyPlatforms {
	return s.Platforms
}

func (s *DescribeParentPlatformsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParentPlatformsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeParentPlatformsResponseBody) SetPageCount(v int64) *DescribeParentPlatformsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetPageNum(v int64) *DescribeParentPlatformsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetPageSize(v int64) *DescribeParentPlatformsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetPlatforms(v []*DescribeParentPlatformsResponseBodyPlatforms) *DescribeParentPlatformsResponseBody {
	s.Platforms = v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetRequestId(v string) *DescribeParentPlatformsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) SetTotalCount(v int64) *DescribeParentPlatformsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeParentPlatformsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParentPlatformsResponseBodyPlatforms struct {
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// example:
	//
	// true
	ClientAuth *bool `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	// example:
	//
	// 31010*****317542006
	ClientGbId *string `json:"ClientGbId,omitempty" xml:"ClientGbId,omitempty"`
	// example:
	//
	// 192.168.0.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// admin123
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	// example:
	//
	// 5160
	ClientPort *int64 `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// example:
	//
	// user01
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	// example:
	//
	// 2018-12-10T21:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 31000*****2170123451
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// example:
	//
	// 359*****374-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10.10.10.10
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 5060
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeParentPlatformsResponseBodyPlatforms) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformsResponseBodyPlatforms) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetClientAuth() *bool {
	return s.ClientAuth
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetClientGbId() *string {
	return s.ClientGbId
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetClientPassword() *string {
	return s.ClientPassword
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetClientPort() *int64 {
	return s.ClientPort
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetClientUsername() *string {
	return s.ClientUsername
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetDescription() *string {
	return s.Description
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetGbId() *string {
	return s.GbId
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetId() *string {
	return s.Id
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetIp() *string {
	return s.Ip
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetName() *string {
	return s.Name
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetPort() *int64 {
	return s.Port
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) GetStatus() *string {
	return s.Status
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetAutoStart(v bool) *DescribeParentPlatformsResponseBodyPlatforms {
	s.AutoStart = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientAuth(v bool) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientAuth = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientGbId(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientGbId = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientIp(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientIp = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientPassword(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientPassword = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientPort(v int64) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientPort = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetClientUsername(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.ClientUsername = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetCreatedTime(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.CreatedTime = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetDescription(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Description = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetGbId(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetId(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetIp(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Ip = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetName(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Name = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetPort(v int64) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Port = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetProtocol(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Protocol = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) SetStatus(v string) *DescribeParentPlatformsResponseBodyPlatforms {
	s.Status = &v
	return s
}

func (s *DescribeParentPlatformsResponseBodyPlatforms) Validate() error {
	return dara.Validate(s)
}
