// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsAnalysisResourceQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBNodeCategoryList(v *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetDBNodeCategoryList() *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList
	SetDBNodeClassList(v *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetDBNodeClassList() *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList
	SetDBNodeStorageList(v *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetDBNodeStorageList() *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList
	SetEngineVersionList(v *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetEngineVersionList() *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList
	SetModeList(v *DescribeRdsAnalysisResourceQuotasResponseBodyModeList) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetModeList() *DescribeRdsAnalysisResourceQuotasResponseBodyModeList
	SetRequestId(v string) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetRequestId() *string
	SetStorageTypeList(v *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) *DescribeRdsAnalysisResourceQuotasResponseBody
	GetStorageTypeList() *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList
}

type DescribeRdsAnalysisResourceQuotasResponseBody struct {
	// The editions of the MySQL analytic instances.
	DBNodeCategoryList *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList `json:"DBNodeCategoryList,omitempty" xml:"DBNodeCategoryList,omitempty" type:"Struct"`
	// The instance types of the MySQL analytic instances.
	DBNodeClassList *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList `json:"DBNodeClassList,omitempty" xml:"DBNodeClassList,omitempty" type:"Struct"`
	// The storage sizes of the MySQL analytic instances.
	DBNodeStorageList *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList `json:"DBNodeStorageList,omitempty" xml:"DBNodeStorageList,omitempty" type:"Struct"`
	// The versions of the MySQL analytic instances.
	EngineVersionList *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList `json:"EngineVersionList,omitempty" xml:"EngineVersionList,omitempty" type:"Struct"`
	// The modes of the MySQL analytic instances.
	ModeList *DescribeRdsAnalysisResourceQuotasResponseBodyModeList `json:"ModeList,omitempty" xml:"ModeList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1A31D7FA-1826-5843-8807-D2F715E70CB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage types of the MySQL analytic instances.
	StorageTypeList *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Struct"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetDBNodeCategoryList() *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList {
	return s.DBNodeCategoryList
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetDBNodeClassList() *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList {
	return s.DBNodeClassList
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetDBNodeStorageList() *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList {
	return s.DBNodeStorageList
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetEngineVersionList() *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList {
	return s.EngineVersionList
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetModeList() *DescribeRdsAnalysisResourceQuotasResponseBodyModeList {
	return s.ModeList
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) GetStorageTypeList() *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList {
	return s.StorageTypeList
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetDBNodeCategoryList(v *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.DBNodeCategoryList = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetDBNodeClassList(v *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.DBNodeClassList = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetDBNodeStorageList(v *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.DBNodeStorageList = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetEngineVersionList(v *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.EngineVersionList = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetModeList(v *DescribeRdsAnalysisResourceQuotasResponseBodyModeList) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.ModeList = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetRequestId(v string) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) SetStorageTypeList(v *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) *DescribeRdsAnalysisResourceQuotasResponseBody {
	s.StorageTypeList = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBody) Validate() error {
	if s.DBNodeCategoryList != nil {
		if err := s.DBNodeCategoryList.Validate(); err != nil {
			return err
		}
	}
	if s.DBNodeClassList != nil {
		if err := s.DBNodeClassList.Validate(); err != nil {
			return err
		}
	}
	if s.DBNodeStorageList != nil {
		if err := s.DBNodeStorageList.Validate(); err != nil {
			return err
		}
	}
	if s.EngineVersionList != nil {
		if err := s.EngineVersionList.Validate(); err != nil {
			return err
		}
	}
	if s.ModeList != nil {
		if err := s.ModeList.Validate(); err != nil {
			return err
		}
	}
	if s.StorageTypeList != nil {
		if err := s.StorageTypeList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList struct {
	DBNodeCategory []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory `json:"DBNodeCategory,omitempty" xml:"DBNodeCategory,omitempty" type:"Repeated"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) GetDBNodeCategory() []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory {
	return s.DBNodeCategory
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) SetDBNodeCategory(v []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList {
	s.DBNodeCategory = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryList) Validate() error {
	if s.DBNodeCategory != nil {
		for _, item := range s.DBNodeCategory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory struct {
	// The display value.
	//
	// example:
	//
	// mixed_storage
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The real value.
	//
	// example:
	//
	// mixed_storage
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) GetText() *string {
	return s.Text
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) GetValue() *string {
	return s.Value
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) SetText(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory {
	s.Text = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) SetValue(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory {
	s.Value = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeCategoryListDBNodeCategory) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList struct {
	DBNodeClass []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty" type:"Repeated"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) GetDBNodeClass() []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass {
	return s.DBNodeClass
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) SetDBNodeClass(v []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList {
	s.DBNodeClass = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassList) Validate() error {
	if s.DBNodeClass != nil {
		for _, item := range s.DBNodeClass {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass struct {
	// The display value.
	//
	// example:
	//
	// E32
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The real value.
	//
	// example:
	//
	// E32
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) GetText() *string {
	return s.Text
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) GetValue() *string {
	return s.Value
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) SetText(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass {
	s.Text = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) SetValue(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass {
	s.Value = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeClassListDBNodeClass) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList struct {
	DBNodeStorage []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty" type:"Repeated"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) GetDBNodeStorage() []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage {
	return s.DBNodeStorage
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) SetDBNodeStorage(v []*DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList {
	s.DBNodeStorage = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageList) Validate() error {
	if s.DBNodeStorage != nil {
		for _, item := range s.DBNodeStorage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage struct {
	// The display value.
	//
	// example:
	//
	// 100
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The real value.
	//
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) GetText() *string {
	return s.Text
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) GetValue() *string {
	return s.Value
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) SetText(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage {
	s.Text = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) SetValue(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage {
	s.Value = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyDBNodeStorageListDBNodeStorage) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList struct {
	EngineVersion []*DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) GetEngineVersion() []*DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion {
	return s.EngineVersion
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) SetEngineVersion(v []*DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList {
	s.EngineVersion = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionList) Validate() error {
	if s.EngineVersion != nil {
		for _, item := range s.EngineVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion struct {
	// The display value.
	//
	// example:
	//
	// 3.0
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The real value.
	//
	// example:
	//
	// 3.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) GetText() *string {
	return s.Text
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) GetValue() *string {
	return s.Value
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) SetText(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion {
	s.Text = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) SetValue(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion {
	s.Value = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyEngineVersionListEngineVersion) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsAnalysisResourceQuotasResponseBodyModeList struct {
	Mode []*DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode `json:"Mode,omitempty" xml:"Mode,omitempty" type:"Repeated"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyModeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyModeList) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeList) GetMode() []*DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode {
	return s.Mode
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeList) SetMode(v []*DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) *DescribeRdsAnalysisResourceQuotasResponseBodyModeList {
	s.Mode = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeList) Validate() error {
	if s.Mode != nil {
		for _, item := range s.Mode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode struct {
	// The display value.
	//
	// example:
	//
	// flexible
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The real value.
	//
	// example:
	//
	// flexible
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) GetText() *string {
	return s.Text
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) GetValue() *string {
	return s.Value
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) SetText(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode {
	s.Text = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) SetValue(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode {
	s.Value = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyModeListMode) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList struct {
	StorageType []*DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType `json:"StorageType,omitempty" xml:"StorageType,omitempty" type:"Repeated"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) GetStorageType() []*DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType {
	return s.StorageType
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) SetStorageType(v []*DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList {
	s.StorageType = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeList) Validate() error {
	if s.StorageType != nil {
		for _, item := range s.StorageType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType struct {
	// The display value.
	//
	// example:
	//
	// cloud_essd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The real value.
	//
	// example:
	//
	// cloud_essd
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) GetText() *string {
	return s.Text
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) GetValue() *string {
	return s.Value
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) SetText(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType {
	s.Text = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) SetValue(v string) *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType {
	s.Value = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponseBodyStorageTypeListStorageType) Validate() error {
	return dara.Validate(s)
}
