// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkBaiLianParameters interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v *SinkBaiLianParametersAfter) *SinkBaiLianParameters
	GetAfter() *SinkBaiLianParametersAfter
	SetApplicationType(v string) *SinkBaiLianParameters
	GetApplicationType() *string
	SetBefore(v *SinkBaiLianParametersBefore) *SinkBaiLianParameters
	GetBefore() *SinkBaiLianParametersBefore
	SetContext(v interface{}) *SinkBaiLianParameters
	GetContext() interface{}
	SetExtend(v interface{}) *SinkBaiLianParameters
	GetExtend() interface{}
	SetOffset(v *SinkBaiLianParametersOffset) *SinkBaiLianParameters
	GetOffset() *SinkBaiLianParametersOffset
	SetOp(v *SinkBaiLianParametersOp) *SinkBaiLianParameters
	GetOp() *SinkBaiLianParametersOp
	SetPartition(v *SinkBaiLianParametersPartition) *SinkBaiLianParameters
	GetPartition() *SinkBaiLianParametersPartition
	SetWorkspaceId(v string) *SinkBaiLianParameters
	GetWorkspaceId() *string
}

type SinkBaiLianParameters struct {
	After           *SinkBaiLianParametersAfter     `json:"After,omitempty" xml:"After,omitempty" type:"Struct"`
	ApplicationType *string                         `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Before          *SinkBaiLianParametersBefore    `json:"Before,omitempty" xml:"Before,omitempty" type:"Struct"`
	Context         interface{}                     `json:"Context,omitempty" xml:"Context,omitempty"`
	Extend          interface{}                     `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Offset          *SinkBaiLianParametersOffset    `json:"Offset,omitempty" xml:"Offset,omitempty" type:"Struct"`
	Op              *SinkBaiLianParametersOp        `json:"Op,omitempty" xml:"Op,omitempty" type:"Struct"`
	Partition       *SinkBaiLianParametersPartition `json:"Partition,omitempty" xml:"Partition,omitempty" type:"Struct"`
	WorkspaceId     *string                         `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SinkBaiLianParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkBaiLianParameters) GoString() string {
	return s.String()
}

func (s *SinkBaiLianParameters) GetAfter() *SinkBaiLianParametersAfter {
	return s.After
}

func (s *SinkBaiLianParameters) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *SinkBaiLianParameters) GetBefore() *SinkBaiLianParametersBefore {
	return s.Before
}

func (s *SinkBaiLianParameters) GetContext() interface{} {
	return s.Context
}

func (s *SinkBaiLianParameters) GetExtend() interface{} {
	return s.Extend
}

func (s *SinkBaiLianParameters) GetOffset() *SinkBaiLianParametersOffset {
	return s.Offset
}

func (s *SinkBaiLianParameters) GetOp() *SinkBaiLianParametersOp {
	return s.Op
}

func (s *SinkBaiLianParameters) GetPartition() *SinkBaiLianParametersPartition {
	return s.Partition
}

func (s *SinkBaiLianParameters) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SinkBaiLianParameters) SetAfter(v *SinkBaiLianParametersAfter) *SinkBaiLianParameters {
	s.After = v
	return s
}

func (s *SinkBaiLianParameters) SetApplicationType(v string) *SinkBaiLianParameters {
	s.ApplicationType = &v
	return s
}

func (s *SinkBaiLianParameters) SetBefore(v *SinkBaiLianParametersBefore) *SinkBaiLianParameters {
	s.Before = v
	return s
}

func (s *SinkBaiLianParameters) SetContext(v interface{}) *SinkBaiLianParameters {
	s.Context = v
	return s
}

func (s *SinkBaiLianParameters) SetExtend(v interface{}) *SinkBaiLianParameters {
	s.Extend = v
	return s
}

func (s *SinkBaiLianParameters) SetOffset(v *SinkBaiLianParametersOffset) *SinkBaiLianParameters {
	s.Offset = v
	return s
}

func (s *SinkBaiLianParameters) SetOp(v *SinkBaiLianParametersOp) *SinkBaiLianParameters {
	s.Op = v
	return s
}

func (s *SinkBaiLianParameters) SetPartition(v *SinkBaiLianParametersPartition) *SinkBaiLianParameters {
	s.Partition = v
	return s
}

func (s *SinkBaiLianParameters) SetWorkspaceId(v string) *SinkBaiLianParameters {
	s.WorkspaceId = &v
	return s
}

func (s *SinkBaiLianParameters) Validate() error {
	if s.After != nil {
		if err := s.After.Validate(); err != nil {
			return err
		}
	}
	if s.Before != nil {
		if err := s.Before.Validate(); err != nil {
			return err
		}
	}
	if s.Offset != nil {
		if err := s.Offset.Validate(); err != nil {
			return err
		}
	}
	if s.Op != nil {
		if err := s.Op.Validate(); err != nil {
			return err
		}
	}
	if s.Partition != nil {
		if err := s.Partition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkBaiLianParametersAfter struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkBaiLianParametersAfter) String() string {
	return dara.Prettify(s)
}

func (s SinkBaiLianParametersAfter) GoString() string {
	return s.String()
}

func (s *SinkBaiLianParametersAfter) GetForm() *string {
	return s.Form
}

func (s *SinkBaiLianParametersAfter) GetTemplate() *string {
	return s.Template
}

func (s *SinkBaiLianParametersAfter) GetValue() *string {
	return s.Value
}

func (s *SinkBaiLianParametersAfter) SetForm(v string) *SinkBaiLianParametersAfter {
	s.Form = &v
	return s
}

func (s *SinkBaiLianParametersAfter) SetTemplate(v string) *SinkBaiLianParametersAfter {
	s.Template = &v
	return s
}

func (s *SinkBaiLianParametersAfter) SetValue(v string) *SinkBaiLianParametersAfter {
	s.Value = &v
	return s
}

func (s *SinkBaiLianParametersAfter) Validate() error {
	return dara.Validate(s)
}

type SinkBaiLianParametersBefore struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkBaiLianParametersBefore) String() string {
	return dara.Prettify(s)
}

func (s SinkBaiLianParametersBefore) GoString() string {
	return s.String()
}

func (s *SinkBaiLianParametersBefore) GetForm() *string {
	return s.Form
}

func (s *SinkBaiLianParametersBefore) GetTemplate() *string {
	return s.Template
}

func (s *SinkBaiLianParametersBefore) GetValue() *string {
	return s.Value
}

func (s *SinkBaiLianParametersBefore) SetForm(v string) *SinkBaiLianParametersBefore {
	s.Form = &v
	return s
}

func (s *SinkBaiLianParametersBefore) SetTemplate(v string) *SinkBaiLianParametersBefore {
	s.Template = &v
	return s
}

func (s *SinkBaiLianParametersBefore) SetValue(v string) *SinkBaiLianParametersBefore {
	s.Value = &v
	return s
}

func (s *SinkBaiLianParametersBefore) Validate() error {
	return dara.Validate(s)
}

type SinkBaiLianParametersOffset struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkBaiLianParametersOffset) String() string {
	return dara.Prettify(s)
}

func (s SinkBaiLianParametersOffset) GoString() string {
	return s.String()
}

func (s *SinkBaiLianParametersOffset) GetForm() *string {
	return s.Form
}

func (s *SinkBaiLianParametersOffset) GetTemplate() *string {
	return s.Template
}

func (s *SinkBaiLianParametersOffset) GetValue() *string {
	return s.Value
}

func (s *SinkBaiLianParametersOffset) SetForm(v string) *SinkBaiLianParametersOffset {
	s.Form = &v
	return s
}

func (s *SinkBaiLianParametersOffset) SetTemplate(v string) *SinkBaiLianParametersOffset {
	s.Template = &v
	return s
}

func (s *SinkBaiLianParametersOffset) SetValue(v string) *SinkBaiLianParametersOffset {
	s.Value = &v
	return s
}

func (s *SinkBaiLianParametersOffset) Validate() error {
	return dara.Validate(s)
}

type SinkBaiLianParametersOp struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkBaiLianParametersOp) String() string {
	return dara.Prettify(s)
}

func (s SinkBaiLianParametersOp) GoString() string {
	return s.String()
}

func (s *SinkBaiLianParametersOp) GetForm() *string {
	return s.Form
}

func (s *SinkBaiLianParametersOp) GetTemplate() *string {
	return s.Template
}

func (s *SinkBaiLianParametersOp) GetValue() *string {
	return s.Value
}

func (s *SinkBaiLianParametersOp) SetForm(v string) *SinkBaiLianParametersOp {
	s.Form = &v
	return s
}

func (s *SinkBaiLianParametersOp) SetTemplate(v string) *SinkBaiLianParametersOp {
	s.Template = &v
	return s
}

func (s *SinkBaiLianParametersOp) SetValue(v string) *SinkBaiLianParametersOp {
	s.Value = &v
	return s
}

func (s *SinkBaiLianParametersOp) Validate() error {
	return dara.Validate(s)
}

type SinkBaiLianParametersPartition struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkBaiLianParametersPartition) String() string {
	return dara.Prettify(s)
}

func (s SinkBaiLianParametersPartition) GoString() string {
	return s.String()
}

func (s *SinkBaiLianParametersPartition) GetForm() *string {
	return s.Form
}

func (s *SinkBaiLianParametersPartition) GetTemplate() *string {
	return s.Template
}

func (s *SinkBaiLianParametersPartition) GetValue() *string {
	return s.Value
}

func (s *SinkBaiLianParametersPartition) SetForm(v string) *SinkBaiLianParametersPartition {
	s.Form = &v
	return s
}

func (s *SinkBaiLianParametersPartition) SetTemplate(v string) *SinkBaiLianParametersPartition {
	s.Template = &v
	return s
}

func (s *SinkBaiLianParametersPartition) SetValue(v string) *SinkBaiLianParametersPartition {
	s.Value = &v
	return s
}

func (s *SinkBaiLianParametersPartition) Validate() error {
	return dara.Validate(s)
}
