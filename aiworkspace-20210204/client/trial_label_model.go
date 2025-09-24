// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrialLabel interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *TrialLabel
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *TrialLabel
	GetGmtModifiedTime() *string
	SetKey(v string) *TrialLabel
	GetKey() *string
	SetTrialId(v string) *TrialLabel
	GetTrialId() *string
	SetValue(v string) *TrialLabel
	GetValue() *string
}

type TrialLabel struct {
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// key
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	TrialId *string `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TrialLabel) String() string {
	return dara.Prettify(s)
}

func (s TrialLabel) GoString() string {
	return s.String()
}

func (s *TrialLabel) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *TrialLabel) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *TrialLabel) GetKey() *string {
	return s.Key
}

func (s *TrialLabel) GetTrialId() *string {
	return s.TrialId
}

func (s *TrialLabel) GetValue() *string {
	return s.Value
}

func (s *TrialLabel) SetGmtCreateTime(v string) *TrialLabel {
	s.GmtCreateTime = &v
	return s
}

func (s *TrialLabel) SetGmtModifiedTime(v string) *TrialLabel {
	s.GmtModifiedTime = &v
	return s
}

func (s *TrialLabel) SetKey(v string) *TrialLabel {
	s.Key = &v
	return s
}

func (s *TrialLabel) SetTrialId(v string) *TrialLabel {
	s.TrialId = &v
	return s
}

func (s *TrialLabel) SetValue(v string) *TrialLabel {
	s.Value = &v
	return s
}

func (s *TrialLabel) Validate() error {
	return dara.Validate(s)
}
