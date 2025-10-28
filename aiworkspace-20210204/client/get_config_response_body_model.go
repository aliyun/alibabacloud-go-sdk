// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *GetConfigResponseBody
	GetCategoryName() *string
	SetConfigKey(v string) *GetConfigResponseBody
	GetConfigKey() *string
	SetConfigValue(v string) *GetConfigResponseBody
	GetConfigValue() *string
	SetGmtCreateTime(v string) *GetConfigResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetConfigResponseBody
	GetGmtModifiedTime() *string
	SetLabels(v []*GetConfigResponseBodyLabels) *GetConfigResponseBody
	GetLabels() []*GetConfigResponseBodyLabels
	SetRequestId(v string) *GetConfigResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *GetConfigResponseBody
	GetWorkspaceId() *string
}

type GetConfigResponseBody struct {
	// The category of the configuration item. Valid values:
	//
	// 	- CommonResourceConfig
	//
	// 	- DLCAutoRecycle
	//
	// 	- DLCPriorityConfig
	//
	// 	- DSWPriorityConfig
	//
	// 	- QuotaMaximumDuration
	//
	// 	- CommonTagConfig
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item. Valid values:
	//
	// 	- tempStoragePath: Temporary storage path. This key can be used only when CategoryName is set to CommonResourceConfig.
	//
	// 	- isAutoRecycle: Automatic recycle configuration. This key can be used only when CategoryName is set to DLCAutoRecycle.
	//
	// 	- priorityConfig: Priority configuration. This key can be used only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// 	- quotaMaximumDuration: Maximum run time of DLC jobs for a quota. This key can be used only when CategoryName is set to QuotaMaximumDuration.
	//
	// 	- predefinedTags: Predefined tags of the workspace. Created resources must include tags.
	//
	// example:
	//
	// tempStoragePath
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// oss://***
	ConfigValue     *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The tags of the configuration item.
	Labels []*GetConfigResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A******C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1234******2
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigResponseBody) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetConfigResponseBody) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *GetConfigResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetConfigResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetConfigResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetConfigResponseBody) GetLabels() []*GetConfigResponseBodyLabels {
	return s.Labels
}

func (s *GetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetConfigResponseBody) SetCategoryName(v string) *GetConfigResponseBody {
	s.CategoryName = &v
	return s
}

func (s *GetConfigResponseBody) SetConfigKey(v string) *GetConfigResponseBody {
	s.ConfigKey = &v
	return s
}

func (s *GetConfigResponseBody) SetConfigValue(v string) *GetConfigResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *GetConfigResponseBody) SetGmtCreateTime(v string) *GetConfigResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetConfigResponseBody) SetGmtModifiedTime(v string) *GetConfigResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetConfigResponseBody) SetLabels(v []*GetConfigResponseBodyLabels) *GetConfigResponseBody {
	s.Labels = v
	return s
}

func (s *GetConfigResponseBody) SetRequestId(v string) *GetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigResponseBody) SetWorkspaceId(v string) *GetConfigResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetConfigResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetConfigResponseBodyLabels struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConfigResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetConfigResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetConfigResponseBodyLabels) GetKey() *string {
	return s.Key
}

func (s *GetConfigResponseBodyLabels) GetValue() *string {
	return s.Value
}

func (s *GetConfigResponseBodyLabels) SetKey(v string) *GetConfigResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetConfigResponseBodyLabels) SetValue(v string) *GetConfigResponseBodyLabels {
	s.Value = &v
	return s
}

func (s *GetConfigResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
