// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLindormV2InstanceSecurityGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *GetLindormV2InstanceSecurityGroupsResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetLindormV2InstanceSecurityGroupsResponseBody
	GetRequestId() *string
	SetSgList(v []*GetLindormV2InstanceSecurityGroupsResponseBodySgList) *GetLindormV2InstanceSecurityGroupsResponseBody
	GetSgList() []*GetLindormV2InstanceSecurityGroupsResponseBodySgList
}

type GetLindormV2InstanceSecurityGroupsResponseBody struct {
	AccessDeniedDetail *string                                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SgList             []*GetLindormV2InstanceSecurityGroupsResponseBodySgList `json:"SgList,omitempty" xml:"SgList,omitempty" type:"Repeated"`
}

func (s GetLindormV2InstanceSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) GetSgList() []*GetLindormV2InstanceSecurityGroupsResponseBodySgList {
	return s.SgList
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) SetAccessDeniedDetail(v string) *GetLindormV2InstanceSecurityGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) SetInstanceId(v string) *GetLindormV2InstanceSecurityGroupsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) SetRequestId(v string) *GetLindormV2InstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) SetSgList(v []*GetLindormV2InstanceSecurityGroupsResponseBodySgList) *GetLindormV2InstanceSecurityGroupsResponseBody {
	s.SgList = v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceSecurityGroupsResponseBodySgList struct {
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	IsAvailable     *bool   `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s GetLindormV2InstanceSecurityGroupsResponseBodySgList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceSecurityGroupsResponseBodySgList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) GetIpList() *string {
	return s.IpList
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) SetGmtModified(v string) *GetLindormV2InstanceSecurityGroupsResponseBodySgList {
	s.GmtModified = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) SetIpList(v string) *GetLindormV2InstanceSecurityGroupsResponseBodySgList {
	s.IpList = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) SetIsAvailable(v bool) *GetLindormV2InstanceSecurityGroupsResponseBodySgList {
	s.IsAvailable = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) SetSecurityGroupId(v string) *GetLindormV2InstanceSecurityGroupsResponseBodySgList {
	s.SecurityGroupId = &v
	return s
}

func (s *GetLindormV2InstanceSecurityGroupsResponseBodySgList) Validate() error {
	return dara.Validate(s)
}
