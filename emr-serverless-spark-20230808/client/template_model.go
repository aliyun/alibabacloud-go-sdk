// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *Template
	GetBizId() *string
	SetCreator(v int64) *Template
	GetCreator() *int64
	SetDisplaySparkVersion(v string) *Template
	GetDisplaySparkVersion() *string
	SetFusion(v bool) *Template
	GetFusion() *bool
	SetGmtCreated(v string) *Template
	GetGmtCreated() *string
	SetGmtModified(v string) *Template
	GetGmtModified() *string
	SetIsDefault(v bool) *Template
	GetIsDefault() *bool
	SetModifier(v int64) *Template
	GetModifier() *int64
	SetName(v string) *Template
	GetName() *string
	SetSparkConf(v []*SparkConf) *Template
	GetSparkConf() []*SparkConf
	SetSparkDriverCores(v int32) *Template
	GetSparkDriverCores() *int32
	SetSparkDriverMemory(v int64) *Template
	GetSparkDriverMemory() *int64
	SetSparkExecutorCores(v int32) *Template
	GetSparkExecutorCores() *int32
	SetSparkExecutorMemory(v int64) *Template
	GetSparkExecutorMemory() *int64
	SetSparkLogLevel(v string) *Template
	GetSparkLogLevel() *string
	SetSparkLogPath(v string) *Template
	GetSparkLogPath() *string
	SetSparkVersion(v string) *Template
	GetSparkVersion() *string
	SetTemplateType(v string) *Template
	GetTemplateType() *string
}

type Template struct {
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Creator             *int64  `json:"creator,omitempty" xml:"creator,omitempty"`
	DisplaySparkVersion *string `json:"displaySparkVersion,omitempty" xml:"displaySparkVersion,omitempty"`
	Fusion              *bool   `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IsDefault   *bool   `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// This parameter is required.
	Modifier  *int64       `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name      *string      `json:"name,omitempty" xml:"name,omitempty"`
	SparkConf []*SparkConf `json:"sparkConf,omitempty" xml:"sparkConf,omitempty" type:"Repeated"`
	// This parameter is required.
	SparkDriverCores *int32 `json:"sparkDriverCores,omitempty" xml:"sparkDriverCores,omitempty"`
	// This parameter is required.
	SparkDriverMemory *int64 `json:"sparkDriverMemory,omitempty" xml:"sparkDriverMemory,omitempty"`
	// This parameter is required.
	SparkExecutorCores *int32 `json:"sparkExecutorCores,omitempty" xml:"sparkExecutorCores,omitempty"`
	// This parameter is required.
	SparkExecutorMemory *int64 `json:"sparkExecutorMemory,omitempty" xml:"sparkExecutorMemory,omitempty"`
	// This parameter is required.
	SparkLogLevel *string `json:"sparkLogLevel,omitempty" xml:"sparkLogLevel,omitempty"`
	// This parameter is required.
	SparkLogPath *string `json:"sparkLogPath,omitempty" xml:"sparkLogPath,omitempty"`
	// This parameter is required.
	SparkVersion *string `json:"sparkVersion,omitempty" xml:"sparkVersion,omitempty"`
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s Template) String() string {
	return dara.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) GetBizId() *string {
	return s.BizId
}

func (s *Template) GetCreator() *int64 {
	return s.Creator
}

func (s *Template) GetDisplaySparkVersion() *string {
	return s.DisplaySparkVersion
}

func (s *Template) GetFusion() *bool {
	return s.Fusion
}

func (s *Template) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *Template) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Template) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *Template) GetModifier() *int64 {
	return s.Modifier
}

func (s *Template) GetName() *string {
	return s.Name
}

func (s *Template) GetSparkConf() []*SparkConf {
	return s.SparkConf
}

func (s *Template) GetSparkDriverCores() *int32 {
	return s.SparkDriverCores
}

func (s *Template) GetSparkDriverMemory() *int64 {
	return s.SparkDriverMemory
}

func (s *Template) GetSparkExecutorCores() *int32 {
	return s.SparkExecutorCores
}

func (s *Template) GetSparkExecutorMemory() *int64 {
	return s.SparkExecutorMemory
}

func (s *Template) GetSparkLogLevel() *string {
	return s.SparkLogLevel
}

func (s *Template) GetSparkLogPath() *string {
	return s.SparkLogPath
}

func (s *Template) GetSparkVersion() *string {
	return s.SparkVersion
}

func (s *Template) GetTemplateType() *string {
	return s.TemplateType
}

func (s *Template) SetBizId(v string) *Template {
	s.BizId = &v
	return s
}

func (s *Template) SetCreator(v int64) *Template {
	s.Creator = &v
	return s
}

func (s *Template) SetDisplaySparkVersion(v string) *Template {
	s.DisplaySparkVersion = &v
	return s
}

func (s *Template) SetFusion(v bool) *Template {
	s.Fusion = &v
	return s
}

func (s *Template) SetGmtCreated(v string) *Template {
	s.GmtCreated = &v
	return s
}

func (s *Template) SetGmtModified(v string) *Template {
	s.GmtModified = &v
	return s
}

func (s *Template) SetIsDefault(v bool) *Template {
	s.IsDefault = &v
	return s
}

func (s *Template) SetModifier(v int64) *Template {
	s.Modifier = &v
	return s
}

func (s *Template) SetName(v string) *Template {
	s.Name = &v
	return s
}

func (s *Template) SetSparkConf(v []*SparkConf) *Template {
	s.SparkConf = v
	return s
}

func (s *Template) SetSparkDriverCores(v int32) *Template {
	s.SparkDriverCores = &v
	return s
}

func (s *Template) SetSparkDriverMemory(v int64) *Template {
	s.SparkDriverMemory = &v
	return s
}

func (s *Template) SetSparkExecutorCores(v int32) *Template {
	s.SparkExecutorCores = &v
	return s
}

func (s *Template) SetSparkExecutorMemory(v int64) *Template {
	s.SparkExecutorMemory = &v
	return s
}

func (s *Template) SetSparkLogLevel(v string) *Template {
	s.SparkLogLevel = &v
	return s
}

func (s *Template) SetSparkLogPath(v string) *Template {
	s.SparkLogPath = &v
	return s
}

func (s *Template) SetSparkVersion(v string) *Template {
	s.SparkVersion = &v
	return s
}

func (s *Template) SetTemplateType(v string) *Template {
	s.TemplateType = &v
	return s
}

func (s *Template) Validate() error {
	return dara.Validate(s)
}
