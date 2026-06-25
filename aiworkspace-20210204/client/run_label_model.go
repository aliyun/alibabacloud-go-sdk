// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLabel interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *RunLabel
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *RunLabel
	GetGmtModifiedTime() *string
	SetKey(v string) *RunLabel
	GetKey() *string
	SetRunId(v string) *RunLabel
	GetRunId() *string
	SetValue(v string) *RunLabel
	GetValue() *string
}

type RunLabel struct {
	// Time when the label was created
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Time when the label was last modified
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Key of the label
	//
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// ID of the Run that owns this label
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// Value of the label
	//
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunLabel) String() string {
	return dara.Prettify(s)
}

func (s RunLabel) GoString() string {
	return s.String()
}

func (s *RunLabel) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *RunLabel) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *RunLabel) GetKey() *string {
	return s.Key
}

func (s *RunLabel) GetRunId() *string {
	return s.RunId
}

func (s *RunLabel) GetValue() *string {
	return s.Value
}

func (s *RunLabel) SetGmtCreateTime(v string) *RunLabel {
	s.GmtCreateTime = &v
	return s
}

func (s *RunLabel) SetGmtModifiedTime(v string) *RunLabel {
	s.GmtModifiedTime = &v
	return s
}

func (s *RunLabel) SetKey(v string) *RunLabel {
	s.Key = &v
	return s
}

func (s *RunLabel) SetRunId(v string) *RunLabel {
	s.RunId = &v
	return s
}

func (s *RunLabel) SetValue(v string) *RunLabel {
	s.Value = &v
	return s
}

func (s *RunLabel) Validate() error {
	return dara.Validate(s)
}
