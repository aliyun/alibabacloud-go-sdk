// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillHubConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetSkillHubConfigResponseBody
	GetCreateTime() *string
	SetOssBucketName(v string) *GetSkillHubConfigResponseBody
	GetOssBucketName() *string
	SetOssRegionId(v string) *GetSkillHubConfigResponseBody
	GetOssRegionId() *string
	SetRequestId(v string) *GetSkillHubConfigResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *GetSkillHubConfigResponseBody
	GetUpdateTime() *string
}

type GetSkillHubConfigResponseBody struct {
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// mybucket
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2849EE73-AFFA-5AFD-9575-12FA886451DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSkillHubConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillHubConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillHubConfigResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSkillHubConfigResponseBody) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetSkillHubConfigResponseBody) GetOssRegionId() *string {
	return s.OssRegionId
}

func (s *GetSkillHubConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillHubConfigResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSkillHubConfigResponseBody) SetCreateTime(v string) *GetSkillHubConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSkillHubConfigResponseBody) SetOssBucketName(v string) *GetSkillHubConfigResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *GetSkillHubConfigResponseBody) SetOssRegionId(v string) *GetSkillHubConfigResponseBody {
	s.OssRegionId = &v
	return s
}

func (s *GetSkillHubConfigResponseBody) SetRequestId(v string) *GetSkillHubConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillHubConfigResponseBody) SetUpdateTime(v string) *GetSkillHubConfigResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillHubConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
