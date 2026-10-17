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
	// The creation time.
	//
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// is_evaluation
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The associated Run ID.
	//
	// example:
	//
	// run-1qJhzJ2YXgX****
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
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
