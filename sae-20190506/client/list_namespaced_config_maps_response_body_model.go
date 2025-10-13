// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacedConfigMapsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNamespacedConfigMapsResponseBody
	GetCode() *string
	SetData(v *ListNamespacedConfigMapsResponseBodyData) *ListNamespacedConfigMapsResponseBody
	GetData() *ListNamespacedConfigMapsResponseBodyData
	SetErrorCode(v string) *ListNamespacedConfigMapsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListNamespacedConfigMapsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNamespacedConfigMapsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNamespacedConfigMapsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListNamespacedConfigMapsResponseBody
	GetTraceId() *string
}

type ListNamespacedConfigMapsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the call was successful.
	//
	// 	- **3xx**: indicates that the call was redirected.
	//
	// 	- **4xx**: indicates that the call failed.
	//
	// 	- **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *ListNamespacedConfigMapsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error code. Valid values:
	//
	// - If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned information. Valid values:
	//
	// 	- If the call is successful, **success*	- is returned.
	//
	// 	- If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the ConfigMap instances were obtained. Valid values:
	//
	// 	- **true**: The instances were obtained.
	//
	// 	- **false**: The instances failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListNamespacedConfigMapsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNamespacedConfigMapsResponseBody) GetData() *ListNamespacedConfigMapsResponseBodyData {
	return s.Data
}

func (s *ListNamespacedConfigMapsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNamespacedConfigMapsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNamespacedConfigMapsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacedConfigMapsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNamespacedConfigMapsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListNamespacedConfigMapsResponseBody) SetCode(v string) *ListNamespacedConfigMapsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetData(v *ListNamespacedConfigMapsResponseBodyData) *ListNamespacedConfigMapsResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetErrorCode(v string) *ListNamespacedConfigMapsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetMessage(v string) *ListNamespacedConfigMapsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetRequestId(v string) *ListNamespacedConfigMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetSuccess(v bool) *ListNamespacedConfigMapsResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetTraceId(v string) *ListNamespacedConfigMapsResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespacedConfigMapsResponseBodyData struct {
	// The ConfigMap instances.
	ConfigMaps []*ListNamespacedConfigMapsResponseBodyDataConfigMaps `json:"ConfigMaps,omitempty" xml:"ConfigMaps,omitempty" type:"Repeated"`
}

func (s ListNamespacedConfigMapsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBodyData) GetConfigMaps() []*ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	return s.ConfigMaps
}

func (s *ListNamespacedConfigMapsResponseBodyData) SetConfigMaps(v []*ListNamespacedConfigMapsResponseBodyDataConfigMaps) *ListNamespacedConfigMapsResponseBodyData {
	s.ConfigMaps = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyData) Validate() error {
	if s.ConfigMaps != nil {
		for _, item := range s.ConfigMaps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespacedConfigMapsResponseBodyDataConfigMaps struct {
	// The ID of the ConfigMap instance.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1593760185111
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data of ConfigMap key-value pairs. Format:
	//
	// {"k1":"v1", "k2":"v2"}
	//
	// k specifies a key and v specifies a value. For more information, see [Manage and use configurations](https://help.aliyun.com/document_detail/171326.html).
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test-desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the ConfigMap instance.
	//
	// example:
	//
	// test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The application that is associated with the instance.
	RelateApps []*ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps `json:"RelateApps,omitempty" xml:"RelateApps,omitempty" type:"Repeated"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1593760185111
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMaps) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMaps) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetData() map[string]interface{} {
	return s.Data
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetDescription() *string {
	return s.Description
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetName() *string {
	return s.Name
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetRelateApps() []*ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps {
	return s.RelateApps
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetConfigMapId(v int64) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.ConfigMapId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetCreateTime(v int64) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.CreateTime = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetData(v map[string]interface{}) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.Data = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetDescription(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.Description = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetName(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.Name = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetNamespaceId(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetRelateApps(v []*ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.RelateApps = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetUpdateTime(v int64) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.UpdateTime = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) Validate() error {
	if s.RelateApps != nil {
		for _, item := range s.RelateApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps struct {
	// The ID of the application.
	//
	// example:
	//
	// f16b4000-9058-4c22-a49d-49a28f0b****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) GetAppId() *string {
	return s.AppId
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) GetAppName() *string {
	return s.AppName
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) SetAppId(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps {
	s.AppId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) SetAppName(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps {
	s.AppName = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) Validate() error {
	return dara.Validate(s)
}
