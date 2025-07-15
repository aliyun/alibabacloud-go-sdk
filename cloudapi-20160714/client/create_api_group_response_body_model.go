// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasePath(v string) *CreateApiGroupResponseBody
	GetBasePath() *string
	SetDescription(v string) *CreateApiGroupResponseBody
	GetDescription() *string
	SetGroupId(v string) *CreateApiGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *CreateApiGroupResponseBody
	GetGroupName() *string
	SetInstanceId(v string) *CreateApiGroupResponseBody
	GetInstanceId() *string
	SetInstanceType(v string) *CreateApiGroupResponseBody
	GetInstanceType() *string
	SetRequestId(v string) *CreateApiGroupResponseBody
	GetRequestId() *string
	SetSubDomain(v string) *CreateApiGroupResponseBody
	GetSubDomain() *string
	SetTagStatus(v bool) *CreateApiGroupResponseBody
	GetTagStatus() *bool
}

type CreateApiGroupResponseBody struct {
	// example:
	//
	// /qqq
	BasePath *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	// example:
	//
	// The weather informations
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 523e8dc7bbe04613b5b1d726c2a7xx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// Weather
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// apigateway-cn-v6419k43xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// VPC_SHARED
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// FF3B7D81-66AE-47E0-BF69-157DCF187514
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 523e8dc7bbe04613b5b1d726xxxxxxxx-cn-hangzhou.alicloudapi.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// example:
	//
	// True
	TagStatus *bool `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
}

func (s CreateApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiGroupResponseBody) GetBasePath() *string {
	return s.BasePath
}

func (s *CreateApiGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateApiGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateApiGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateApiGroupResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApiGroupResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiGroupResponseBody) GetSubDomain() *string {
	return s.SubDomain
}

func (s *CreateApiGroupResponseBody) GetTagStatus() *bool {
	return s.TagStatus
}

func (s *CreateApiGroupResponseBody) SetBasePath(v string) *CreateApiGroupResponseBody {
	s.BasePath = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetDescription(v string) *CreateApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetGroupId(v string) *CreateApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetGroupName(v string) *CreateApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetInstanceId(v string) *CreateApiGroupResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetInstanceType(v string) *CreateApiGroupResponseBody {
	s.InstanceType = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetRequestId(v string) *CreateApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetSubDomain(v string) *CreateApiGroupResponseBody {
	s.SubDomain = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetTagStatus(v bool) *CreateApiGroupResponseBody {
	s.TagStatus = &v
	return s
}

func (s *CreateApiGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
