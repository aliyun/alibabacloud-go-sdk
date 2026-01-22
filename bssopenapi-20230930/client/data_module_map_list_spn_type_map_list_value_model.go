// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataModuleMapListSpnTypeMapListValue interface {
	dara.Model
	String() string
	GoString() string
	SetFilterModules(v []*DataModuleMapListSpnTypeMapListValueFilterModules) *DataModuleMapListSpnTypeMapListValue
	GetFilterModules() []*DataModuleMapListSpnTypeMapListValueFilterModules
	SetShowModules(v []*DataModuleMapListSpnTypeMapListValueShowModules) *DataModuleMapListSpnTypeMapListValue
	GetShowModules() []*DataModuleMapListSpnTypeMapListValueShowModules
}

type DataModuleMapListSpnTypeMapListValue struct {
	FilterModules []*DataModuleMapListSpnTypeMapListValueFilterModules `json:"FilterModules,omitempty" xml:"FilterModules,omitempty" type:"Repeated"`
	ShowModules   []*DataModuleMapListSpnTypeMapListValueShowModules   `json:"ShowModules,omitempty" xml:"ShowModules,omitempty" type:"Repeated"`
}

func (s DataModuleMapListSpnTypeMapListValue) String() string {
	return dara.Prettify(s)
}

func (s DataModuleMapListSpnTypeMapListValue) GoString() string {
	return s.String()
}

func (s *DataModuleMapListSpnTypeMapListValue) GetFilterModules() []*DataModuleMapListSpnTypeMapListValueFilterModules {
	return s.FilterModules
}

func (s *DataModuleMapListSpnTypeMapListValue) GetShowModules() []*DataModuleMapListSpnTypeMapListValueShowModules {
	return s.ShowModules
}

func (s *DataModuleMapListSpnTypeMapListValue) SetFilterModules(v []*DataModuleMapListSpnTypeMapListValueFilterModules) *DataModuleMapListSpnTypeMapListValue {
	s.FilterModules = v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValue) SetShowModules(v []*DataModuleMapListSpnTypeMapListValueShowModules) *DataModuleMapListSpnTypeMapListValue {
	s.ShowModules = v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValue) Validate() error {
	if s.FilterModules != nil {
		for _, item := range s.FilterModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ShowModules != nil {
		for _, item := range s.ShowModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DataModuleMapListSpnTypeMapListValueFilterModules struct {
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DataModuleMapListSpnTypeMapListValueFilterModules) String() string {
	return dara.Prettify(s)
}

func (s DataModuleMapListSpnTypeMapListValueFilterModules) GoString() string {
	return s.String()
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) SetModuleId(v int64) *DataModuleMapListSpnTypeMapListValueFilterModules {
	s.ModuleId = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) SetModuleCode(v string) *DataModuleMapListSpnTypeMapListValueFilterModules {
	s.ModuleCode = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) SetModuleName(v string) *DataModuleMapListSpnTypeMapListValueFilterModules {
	s.ModuleName = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) Validate() error {
	return dara.Validate(s)
}

type DataModuleMapListSpnTypeMapListValueShowModules struct {
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DataModuleMapListSpnTypeMapListValueShowModules) String() string {
	return dara.Prettify(s)
}

func (s DataModuleMapListSpnTypeMapListValueShowModules) GoString() string {
	return s.String()
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) SetModuleId(v int64) *DataModuleMapListSpnTypeMapListValueShowModules {
	s.ModuleId = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) SetModuleCode(v string) *DataModuleMapListSpnTypeMapListValueShowModules {
	s.ModuleCode = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) SetModuleName(v string) *DataModuleMapListSpnTypeMapListValueShowModules {
	s.ModuleName = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) Validate() error {
	return dara.Validate(s)
}
