// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStackInfoExtInfo interface {
	dara.Model
	String() string
	GoString() string
	SetInfo(v string) *StackInfoExtInfo
	GetInfo() *string
	SetType(v string) *StackInfoExtInfo
	GetType() *string
}

type StackInfoExtInfo struct {
	Info *string `json:"info,omitempty" xml:"info,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StackInfoExtInfo) String() string {
	return dara.Prettify(s)
}

func (s StackInfoExtInfo) GoString() string {
	return s.String()
}

func (s *StackInfoExtInfo) GetInfo() *string {
	return s.Info
}

func (s *StackInfoExtInfo) GetType() *string {
	return s.Type
}

func (s *StackInfoExtInfo) SetInfo(v string) *StackInfoExtInfo {
	s.Info = &v
	return s
}

func (s *StackInfoExtInfo) SetType(v string) *StackInfoExtInfo {
	s.Type = &v
	return s
}

func (s *StackInfoExtInfo) Validate() error {
	return dara.Validate(s)
}
