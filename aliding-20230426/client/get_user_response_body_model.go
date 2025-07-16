// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *GetUserResponseBody
	GetActive() *bool
	SetAdmin(v bool) *GetUserResponseBody
	GetAdmin() *bool
	SetAvatar(v string) *GetUserResponseBody
	GetAvatar() *string
	SetBoss(v bool) *GetUserResponseBody
	GetBoss() *bool
	SetDeptIdList(v []*int64) *GetUserResponseBody
	GetDeptIdList() []*int64
	SetDeptOrderList(v []*GetUserResponseBodyDeptOrderList) *GetUserResponseBody
	GetDeptOrderList() []*GetUserResponseBodyDeptOrderList
	SetEmail(v string) *GetUserResponseBody
	GetEmail() *string
	SetExclusiveAccount(v bool) *GetUserResponseBody
	GetExclusiveAccount() *bool
	SetExclusiveAccountCorpId(v string) *GetUserResponseBody
	GetExclusiveAccountCorpId() *string
	SetExclusiveAccountCorpName(v string) *GetUserResponseBody
	GetExclusiveAccountCorpName() *string
	SetExclusiveAccountType(v string) *GetUserResponseBody
	GetExclusiveAccountType() *string
	SetExtension(v string) *GetUserResponseBody
	GetExtension() *string
	SetHideMobile(v bool) *GetUserResponseBody
	GetHideMobile() *bool
	SetHiredDate(v int64) *GetUserResponseBody
	GetHiredDate() *int64
	SetJobNumber(v string) *GetUserResponseBody
	GetJobNumber() *string
	SetLeaderInDept(v []*GetUserResponseBodyLeaderInDept) *GetUserResponseBody
	GetLeaderInDept() []*GetUserResponseBodyLeaderInDept
	SetLoginId(v string) *GetUserResponseBody
	GetLoginId() *string
	SetManagerUserid(v string) *GetUserResponseBody
	GetManagerUserid() *string
	SetMobile(v string) *GetUserResponseBody
	GetMobile() *string
	SetName(v string) *GetUserResponseBody
	GetName() *string
	SetNickname(v string) *GetUserResponseBody
	GetNickname() *string
	SetOrgEmail(v string) *GetUserResponseBody
	GetOrgEmail() *string
	SetRealAuthed(v bool) *GetUserResponseBody
	GetRealAuthed() *bool
	SetRemark(v string) *GetUserResponseBody
	GetRemark() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetRoleList(v []*GetUserResponseBodyRoleList) *GetUserResponseBody
	GetRoleList() []*GetUserResponseBodyRoleList
	SetSenior(v bool) *GetUserResponseBody
	GetSenior() *bool
	SetStateCode(v string) *GetUserResponseBody
	GetStateCode() *string
	SetTelephone(v string) *GetUserResponseBody
	GetTelephone() *string
	SetTitle(v string) *GetUserResponseBody
	GetTitle() *string
	SetUnionEmpExt(v *GetUserResponseBodyUnionEmpExt) *GetUserResponseBody
	GetUnionEmpExt() *GetUserResponseBodyUnionEmpExt
	SetUnionid(v string) *GetUserResponseBody
	GetUnionid() *string
	SetUserid(v string) *GetUserResponseBody
	GetUserid() *string
	SetWorkPlace(v string) *GetUserResponseBody
	GetWorkPlace() *string
}

type GetUserResponseBody struct {
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// example:
	//
	// true
	Admin *bool `json:"admin,omitempty" xml:"admin,omitempty"`
	// example:
	//
	// xxx
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// true
	Boss *bool `json:"boss,omitempty" xml:"boss,omitempty"`
	// example:
	//
	// [2,3,4]
	DeptIdList    []*int64                            `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	DeptOrderList []*GetUserResponseBodyDeptOrderList `json:"deptOrderList,omitempty" xml:"deptOrderList,omitempty" type:"Repeated"`
	// example:
	//
	// test@xxx.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// true
	ExclusiveAccount *bool `json:"exclusiveAccount,omitempty" xml:"exclusiveAccount,omitempty"`
	// example:
	//
	// dingxxx
	ExclusiveAccountCorpId *string `json:"exclusiveAccountCorpId,omitempty" xml:"exclusiveAccountCorpId,omitempty"`
	// example:
	//
	// 组织名称
	ExclusiveAccountCorpName *string `json:"exclusiveAccountCorpName,omitempty" xml:"exclusiveAccountCorpName,omitempty"`
	// example:
	//
	// dingtalk
	ExclusiveAccountType *string `json:"exclusiveAccountType,omitempty" xml:"exclusiveAccountType,omitempty"`
	// example:
	//
	// {"爱好":"旅游","年龄":"24"}
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// false
	HideMobile *bool `json:"hideMobile,omitempty" xml:"hideMobile,omitempty"`
	// example:
	//
	// 1597573616828
	HiredDate *int64 `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// example:
	//
	// 4
	JobNumber    *string                            `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	LeaderInDept []*GetUserResponseBodyLeaderInDept `json:"leaderInDept,omitempty" xml:"leaderInDept,omitempty" type:"Repeated"`
	// example:
	//
	// login_id3
	LoginId *string `json:"loginId,omitempty" xml:"loginId,omitempty"`
	// example:
	//
	// manager240
	ManagerUserid *string `json:"managerUserid,omitempty" xml:"managerUserid,omitempty"`
	// example:
	//
	// 18513027676
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 昵称
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// example:
	//
	// test@xxx.com
	OrgEmail *string `json:"orgEmail,omitempty" xml:"orgEmail,omitempty"`
	// example:
	//
	// true
	RealAuthed *bool `json:"realAuthed,omitempty" xml:"realAuthed,omitempty"`
	// example:
	//
	// 备注备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RoleList  []*GetUserResponseBodyRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Senior *bool `json:"senior,omitempty" xml:"senior,omitempty"`
	// example:
	//
	// 86
	StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	// example:
	//
	// 010-86123456-2345
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// example:
	//
	// 技术总监
	Title       *string                         `json:"title,omitempty" xml:"title,omitempty"`
	UnionEmpExt *GetUserResponseBodyUnionEmpExt `json:"unionEmpExt,omitempty" xml:"unionEmpExt,omitempty" type:"Struct"`
	Unionid     *string                         `json:"unionid,omitempty" xml:"unionid,omitempty"`
	// example:
	//
	// zhangsan
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
	// example:
	//
	// 未来park
	WorkPlace *string `json:"workPlace,omitempty" xml:"workPlace,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetActive() *bool {
	return s.Active
}

func (s *GetUserResponseBody) GetAdmin() *bool {
	return s.Admin
}

func (s *GetUserResponseBody) GetAvatar() *string {
	return s.Avatar
}

func (s *GetUserResponseBody) GetBoss() *bool {
	return s.Boss
}

func (s *GetUserResponseBody) GetDeptIdList() []*int64 {
	return s.DeptIdList
}

func (s *GetUserResponseBody) GetDeptOrderList() []*GetUserResponseBodyDeptOrderList {
	return s.DeptOrderList
}

func (s *GetUserResponseBody) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBody) GetExclusiveAccount() *bool {
	return s.ExclusiveAccount
}

func (s *GetUserResponseBody) GetExclusiveAccountCorpId() *string {
	return s.ExclusiveAccountCorpId
}

func (s *GetUserResponseBody) GetExclusiveAccountCorpName() *string {
	return s.ExclusiveAccountCorpName
}

func (s *GetUserResponseBody) GetExclusiveAccountType() *string {
	return s.ExclusiveAccountType
}

func (s *GetUserResponseBody) GetExtension() *string {
	return s.Extension
}

func (s *GetUserResponseBody) GetHideMobile() *bool {
	return s.HideMobile
}

func (s *GetUserResponseBody) GetHiredDate() *int64 {
	return s.HiredDate
}

func (s *GetUserResponseBody) GetJobNumber() *string {
	return s.JobNumber
}

func (s *GetUserResponseBody) GetLeaderInDept() []*GetUserResponseBodyLeaderInDept {
	return s.LeaderInDept
}

func (s *GetUserResponseBody) GetLoginId() *string {
	return s.LoginId
}

func (s *GetUserResponseBody) GetManagerUserid() *string {
	return s.ManagerUserid
}

func (s *GetUserResponseBody) GetMobile() *string {
	return s.Mobile
}

func (s *GetUserResponseBody) GetName() *string {
	return s.Name
}

func (s *GetUserResponseBody) GetNickname() *string {
	return s.Nickname
}

func (s *GetUserResponseBody) GetOrgEmail() *string {
	return s.OrgEmail
}

func (s *GetUserResponseBody) GetRealAuthed() *bool {
	return s.RealAuthed
}

func (s *GetUserResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetRoleList() []*GetUserResponseBodyRoleList {
	return s.RoleList
}

func (s *GetUserResponseBody) GetSenior() *bool {
	return s.Senior
}

func (s *GetUserResponseBody) GetStateCode() *string {
	return s.StateCode
}

func (s *GetUserResponseBody) GetTelephone() *string {
	return s.Telephone
}

func (s *GetUserResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetUserResponseBody) GetUnionEmpExt() *GetUserResponseBodyUnionEmpExt {
	return s.UnionEmpExt
}

func (s *GetUserResponseBody) GetUnionid() *string {
	return s.Unionid
}

func (s *GetUserResponseBody) GetUserid() *string {
	return s.Userid
}

func (s *GetUserResponseBody) GetWorkPlace() *string {
	return s.WorkPlace
}

func (s *GetUserResponseBody) SetActive(v bool) *GetUserResponseBody {
	s.Active = &v
	return s
}

func (s *GetUserResponseBody) SetAdmin(v bool) *GetUserResponseBody {
	s.Admin = &v
	return s
}

func (s *GetUserResponseBody) SetAvatar(v string) *GetUserResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetUserResponseBody) SetBoss(v bool) *GetUserResponseBody {
	s.Boss = &v
	return s
}

func (s *GetUserResponseBody) SetDeptIdList(v []*int64) *GetUserResponseBody {
	s.DeptIdList = v
	return s
}

func (s *GetUserResponseBody) SetDeptOrderList(v []*GetUserResponseBodyDeptOrderList) *GetUserResponseBody {
	s.DeptOrderList = v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccount(v bool) *GetUserResponseBody {
	s.ExclusiveAccount = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountCorpId(v string) *GetUserResponseBody {
	s.ExclusiveAccountCorpId = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountCorpName(v string) *GetUserResponseBody {
	s.ExclusiveAccountCorpName = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountType(v string) *GetUserResponseBody {
	s.ExclusiveAccountType = &v
	return s
}

func (s *GetUserResponseBody) SetExtension(v string) *GetUserResponseBody {
	s.Extension = &v
	return s
}

func (s *GetUserResponseBody) SetHideMobile(v bool) *GetUserResponseBody {
	s.HideMobile = &v
	return s
}

func (s *GetUserResponseBody) SetHiredDate(v int64) *GetUserResponseBody {
	s.HiredDate = &v
	return s
}

func (s *GetUserResponseBody) SetJobNumber(v string) *GetUserResponseBody {
	s.JobNumber = &v
	return s
}

func (s *GetUserResponseBody) SetLeaderInDept(v []*GetUserResponseBodyLeaderInDept) *GetUserResponseBody {
	s.LeaderInDept = v
	return s
}

func (s *GetUserResponseBody) SetLoginId(v string) *GetUserResponseBody {
	s.LoginId = &v
	return s
}

func (s *GetUserResponseBody) SetManagerUserid(v string) *GetUserResponseBody {
	s.ManagerUserid = &v
	return s
}

func (s *GetUserResponseBody) SetMobile(v string) *GetUserResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBody) SetName(v string) *GetUserResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserResponseBody) SetNickname(v string) *GetUserResponseBody {
	s.Nickname = &v
	return s
}

func (s *GetUserResponseBody) SetOrgEmail(v string) *GetUserResponseBody {
	s.OrgEmail = &v
	return s
}

func (s *GetUserResponseBody) SetRealAuthed(v bool) *GetUserResponseBody {
	s.RealAuthed = &v
	return s
}

func (s *GetUserResponseBody) SetRemark(v string) *GetUserResponseBody {
	s.Remark = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetRoleList(v []*GetUserResponseBodyRoleList) *GetUserResponseBody {
	s.RoleList = v
	return s
}

func (s *GetUserResponseBody) SetSenior(v bool) *GetUserResponseBody {
	s.Senior = &v
	return s
}

func (s *GetUserResponseBody) SetStateCode(v string) *GetUserResponseBody {
	s.StateCode = &v
	return s
}

func (s *GetUserResponseBody) SetTelephone(v string) *GetUserResponseBody {
	s.Telephone = &v
	return s
}

func (s *GetUserResponseBody) SetTitle(v string) *GetUserResponseBody {
	s.Title = &v
	return s
}

func (s *GetUserResponseBody) SetUnionEmpExt(v *GetUserResponseBodyUnionEmpExt) *GetUserResponseBody {
	s.UnionEmpExt = v
	return s
}

func (s *GetUserResponseBody) SetUnionid(v string) *GetUserResponseBody {
	s.Unionid = &v
	return s
}

func (s *GetUserResponseBody) SetUserid(v string) *GetUserResponseBody {
	s.Userid = &v
	return s
}

func (s *GetUserResponseBody) SetWorkPlace(v string) *GetUserResponseBody {
	s.WorkPlace = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyDeptOrderList struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Order  *int64 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetUserResponseBodyDeptOrderList) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyDeptOrderList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDeptOrderList) GetDeptId() *int64 {
	return s.DeptId
}

func (s *GetUserResponseBodyDeptOrderList) GetOrder() *int64 {
	return s.Order
}

func (s *GetUserResponseBodyDeptOrderList) SetDeptId(v int64) *GetUserResponseBodyDeptOrderList {
	s.DeptId = &v
	return s
}

func (s *GetUserResponseBodyDeptOrderList) SetOrder(v int64) *GetUserResponseBodyDeptOrderList {
	s.Order = &v
	return s
}

func (s *GetUserResponseBodyDeptOrderList) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyLeaderInDept struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Leader *bool  `json:"leader,omitempty" xml:"leader,omitempty"`
}

func (s GetUserResponseBodyLeaderInDept) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyLeaderInDept) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyLeaderInDept) GetDeptId() *int64 {
	return s.DeptId
}

func (s *GetUserResponseBodyLeaderInDept) GetLeader() *bool {
	return s.Leader
}

func (s *GetUserResponseBodyLeaderInDept) SetDeptId(v int64) *GetUserResponseBodyLeaderInDept {
	s.DeptId = &v
	return s
}

func (s *GetUserResponseBodyLeaderInDept) SetLeader(v bool) *GetUserResponseBodyLeaderInDept {
	s.Leader = &v
	return s
}

func (s *GetUserResponseBodyLeaderInDept) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyRoleList struct {
	// example:
	//
	// 职务
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetUserResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyRoleList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetUserResponseBodyRoleList) GetId() *int64 {
	return s.Id
}

func (s *GetUserResponseBodyRoleList) GetName() *string {
	return s.Name
}

func (s *GetUserResponseBodyRoleList) SetGroupName(v string) *GetUserResponseBodyRoleList {
	s.GroupName = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetId(v int64) *GetUserResponseBodyRoleList {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetName(v string) *GetUserResponseBodyRoleList {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyRoleList) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUnionEmpExt struct {
	// example:
	//
	// dingxxx
	CorpId          *string                                          `json:"corpId,omitempty" xml:"corpId,omitempty"`
	UnionEmpMapList []*GetUserResponseBodyUnionEmpExtUnionEmpMapList `json:"unionEmpMapList,omitempty" xml:"unionEmpMapList,omitempty" type:"Repeated"`
	// example:
	//
	// zhangsan
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetUserResponseBodyUnionEmpExt) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExt) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExt) GetCorpId() *string {
	return s.CorpId
}

func (s *GetUserResponseBodyUnionEmpExt) GetUnionEmpMapList() []*GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	return s.UnionEmpMapList
}

func (s *GetUserResponseBodyUnionEmpExt) GetUserid() *string {
	return s.Userid
}

func (s *GetUserResponseBodyUnionEmpExt) SetCorpId(v string) *GetUserResponseBodyUnionEmpExt {
	s.CorpId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetUnionEmpMapList(v []*GetUserResponseBodyUnionEmpExtUnionEmpMapList) *GetUserResponseBodyUnionEmpExt {
	s.UnionEmpMapList = v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetUserid(v string) *GetUserResponseBodyUnionEmpExt {
	s.Userid = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUnionEmpExtUnionEmpMapList struct {
	// example:
	//
	// dingxxx
	CropId *string `json:"cropId,omitempty" xml:"cropId,omitempty"`
	// example:
	//
	// zhangsan
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) GetCropId() *string {
	return s.CropId
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) GetUserid() *string {
	return s.Userid
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetCropId(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.CropId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetUserid(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.Userid = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) Validate() error {
	return dara.Validate(s)
}
