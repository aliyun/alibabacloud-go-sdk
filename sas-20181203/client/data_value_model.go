// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetCveNum(v int32) *DataValue
	GetCveNum() *int32
	SetEmgNum(v int32) *DataValue
	GetEmgNum() *int32
	SetSysNum(v int32) *DataValue
	GetSysNum() *int32
	SetCmsNum(v int32) *DataValue
	GetCmsNum() *int32
	SetAppNum(v int32) *DataValue
	GetAppNum() *int32
	SetScaNum(v int32) *DataValue
	GetScaNum() *int32
	SetVulAsapSum(v int32) *DataValue
	GetVulAsapSum() *int32
	SetVulLaterSum(v int32) *DataValue
	GetVulLaterSum() *int32
	SetVulNntfSum(v int32) *DataValue
	GetVulNntfSum() *int32
	SetSysAsapNum(v int32) *DataValue
	GetSysAsapNum() *int32
}

type DataValue struct {
	// example:
	//
	// 1
	CveNum *int32 `json:"CveNum,omitempty" xml:"CveNum,omitempty"`
	// example:
	//
	// 0
	EmgNum *int32 `json:"EmgNum,omitempty" xml:"EmgNum,omitempty"`
	// example:
	//
	// 0
	SysNum *int32 `json:"SysNum,omitempty" xml:"SysNum,omitempty"`
	// example:
	//
	// 0
	CmsNum *int32 `json:"CmsNum,omitempty" xml:"CmsNum,omitempty"`
	// example:
	//
	// 0
	AppNum *int32 `json:"AppNum,omitempty" xml:"AppNum,omitempty"`
	// example:
	//
	// 2
	ScaNum *int32 `json:"ScaNum,omitempty" xml:"ScaNum,omitempty"`
	// example:
	//
	// 1
	VulAsapSum *int32 `json:"VulAsapSum,omitempty" xml:"VulAsapSum,omitempty"`
	// example:
	//
	// 1
	VulLaterSum *int32 `json:"VulLaterSum,omitempty" xml:"VulLaterSum,omitempty"`
	// example:
	//
	// 1
	VulNntfSum *int32 `json:"VulNntfSum,omitempty" xml:"VulNntfSum,omitempty"`
	// example:
	//
	// 1
	SysAsapNum *int32 `json:"SysAsapNum,omitempty" xml:"SysAsapNum,omitempty"`
}

func (s DataValue) String() string {
	return dara.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) GetCveNum() *int32 {
	return s.CveNum
}

func (s *DataValue) GetEmgNum() *int32 {
	return s.EmgNum
}

func (s *DataValue) GetSysNum() *int32 {
	return s.SysNum
}

func (s *DataValue) GetCmsNum() *int32 {
	return s.CmsNum
}

func (s *DataValue) GetAppNum() *int32 {
	return s.AppNum
}

func (s *DataValue) GetScaNum() *int32 {
	return s.ScaNum
}

func (s *DataValue) GetVulAsapSum() *int32 {
	return s.VulAsapSum
}

func (s *DataValue) GetVulLaterSum() *int32 {
	return s.VulLaterSum
}

func (s *DataValue) GetVulNntfSum() *int32 {
	return s.VulNntfSum
}

func (s *DataValue) GetSysAsapNum() *int32 {
	return s.SysAsapNum
}

func (s *DataValue) SetCveNum(v int32) *DataValue {
	s.CveNum = &v
	return s
}

func (s *DataValue) SetEmgNum(v int32) *DataValue {
	s.EmgNum = &v
	return s
}

func (s *DataValue) SetSysNum(v int32) *DataValue {
	s.SysNum = &v
	return s
}

func (s *DataValue) SetCmsNum(v int32) *DataValue {
	s.CmsNum = &v
	return s
}

func (s *DataValue) SetAppNum(v int32) *DataValue {
	s.AppNum = &v
	return s
}

func (s *DataValue) SetScaNum(v int32) *DataValue {
	s.ScaNum = &v
	return s
}

func (s *DataValue) SetVulAsapSum(v int32) *DataValue {
	s.VulAsapSum = &v
	return s
}

func (s *DataValue) SetVulLaterSum(v int32) *DataValue {
	s.VulLaterSum = &v
	return s
}

func (s *DataValue) SetVulNntfSum(v int32) *DataValue {
	s.VulNntfSum = &v
	return s
}

func (s *DataValue) SetSysAsapNum(v int32) *DataValue {
	s.SysAsapNum = &v
	return s
}

func (s *DataValue) Validate() error {
	return dara.Validate(s)
}
