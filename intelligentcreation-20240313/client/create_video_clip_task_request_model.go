// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoClipTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunMainId(v string) *CreateVideoClipTaskRequest
	GetAliyunMainId() *string
	SetDescription(v string) *CreateVideoClipTaskRequest
	GetDescription() *string
	SetOssKeys(v []*string) *CreateVideoClipTaskRequest
	GetOssKeys() []*string
	SetRequirement(v string) *CreateVideoClipTaskRequest
	GetRequirement() *string
}

type CreateVideoClipTaskRequest struct {
	// example:
	//
	// 1314445556
	AliyunMainId *string   `json:"aliyunMainId,omitempty" xml:"aliyunMainId,omitempty"`
	Description  *string   `json:"description,omitempty" xml:"description,omitempty"`
	OssKeys      []*string `json:"ossKeys,omitempty" xml:"ossKeys,omitempty" type:"Repeated"`
	Requirement  *string   `json:"requirement,omitempty" xml:"requirement,omitempty"`
}

func (s CreateVideoClipTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoClipTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoClipTaskRequest) GetAliyunMainId() *string {
	return s.AliyunMainId
}

func (s *CreateVideoClipTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVideoClipTaskRequest) GetOssKeys() []*string {
	return s.OssKeys
}

func (s *CreateVideoClipTaskRequest) GetRequirement() *string {
	return s.Requirement
}

func (s *CreateVideoClipTaskRequest) SetAliyunMainId(v string) *CreateVideoClipTaskRequest {
	s.AliyunMainId = &v
	return s
}

func (s *CreateVideoClipTaskRequest) SetDescription(v string) *CreateVideoClipTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateVideoClipTaskRequest) SetOssKeys(v []*string) *CreateVideoClipTaskRequest {
	s.OssKeys = v
	return s
}

func (s *CreateVideoClipTaskRequest) SetRequirement(v string) *CreateVideoClipTaskRequest {
	s.Requirement = &v
	return s
}

func (s *CreateVideoClipTaskRequest) Validate() error {
	return dara.Validate(s)
}
