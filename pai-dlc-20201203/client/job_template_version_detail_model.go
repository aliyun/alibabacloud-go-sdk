// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobTemplateVersionDetail interface {
	dara.Model
	String() string
	GoString() string
	SetConstraints(v map[string]*string) *JobTemplateVersionDetail
	GetConstraints() map[string]*string
	SetContent(v string) *JobTemplateVersionDetail
	GetContent() *string
	SetCreatedBy(v string) *JobTemplateVersionDetail
	GetCreatedBy() *string
	SetGmtCreated(v string) *JobTemplateVersionDetail
	GetGmtCreated() *string
	SetVersion(v int32) *JobTemplateVersionDetail
	GetVersion() *int32
}

type JobTemplateVersionDetail struct {
	Constraints map[string]*string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// 任务模板的配置内容，支持 CreateJob 接口的所有参数字段，以 JSON 对象存储
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建该版本的用户ID
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// 该版本的创建时间
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// 模板版本号
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s JobTemplateVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s JobTemplateVersionDetail) GoString() string {
	return s.String()
}

func (s *JobTemplateVersionDetail) GetConstraints() map[string]*string {
	return s.Constraints
}

func (s *JobTemplateVersionDetail) GetContent() *string {
	return s.Content
}

func (s *JobTemplateVersionDetail) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *JobTemplateVersionDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *JobTemplateVersionDetail) GetVersion() *int32 {
	return s.Version
}

func (s *JobTemplateVersionDetail) SetConstraints(v map[string]*string) *JobTemplateVersionDetail {
	s.Constraints = v
	return s
}

func (s *JobTemplateVersionDetail) SetContent(v string) *JobTemplateVersionDetail {
	s.Content = &v
	return s
}

func (s *JobTemplateVersionDetail) SetCreatedBy(v string) *JobTemplateVersionDetail {
	s.CreatedBy = &v
	return s
}

func (s *JobTemplateVersionDetail) SetGmtCreated(v string) *JobTemplateVersionDetail {
	s.GmtCreated = &v
	return s
}

func (s *JobTemplateVersionDetail) SetVersion(v int32) *JobTemplateVersionDetail {
	s.Version = &v
	return s
}

func (s *JobTemplateVersionDetail) Validate() error {
	return dara.Validate(s)
}
