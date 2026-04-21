// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserAPICountInfo interface {
	dara.Model
	String() string
	GoString() string
	SetApiType(v string) *UserAPICountInfo
	GetApiType() *string
	SetScope(v string) *UserAPICountInfo
	GetScope() *string
	SetUsedCount(v int32) *UserAPICountInfo
	GetUsedCount() *int32
}

type UserAPICountInfo struct {
	ApiType   *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	Scope     *string `json:"scope,omitempty" xml:"scope,omitempty"`
	UsedCount *int32  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
}

func (s UserAPICountInfo) String() string {
	return dara.Prettify(s)
}

func (s UserAPICountInfo) GoString() string {
	return s.String()
}

func (s *UserAPICountInfo) GetApiType() *string {
	return s.ApiType
}

func (s *UserAPICountInfo) GetScope() *string {
	return s.Scope
}

func (s *UserAPICountInfo) GetUsedCount() *int32 {
	return s.UsedCount
}

func (s *UserAPICountInfo) SetApiType(v string) *UserAPICountInfo {
	s.ApiType = &v
	return s
}

func (s *UserAPICountInfo) SetScope(v string) *UserAPICountInfo {
	s.Scope = &v
	return s
}

func (s *UserAPICountInfo) SetUsedCount(v int32) *UserAPICountInfo {
	s.UsedCount = &v
	return s
}

func (s *UserAPICountInfo) Validate() error {
	return dara.Validate(s)
}
