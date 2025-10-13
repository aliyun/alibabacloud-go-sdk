// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeConfigMapResponseBody
	GetCode() *string
	SetData(v *DescribeConfigMapResponseBodyData) *DescribeConfigMapResponseBody
	GetData() *DescribeConfigMapResponseBodyData
	SetErrorCode(v string) *DescribeConfigMapResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeConfigMapResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeConfigMapResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeConfigMapResponseBody
	GetTraceId() *string
}

type DescribeConfigMapResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *DescribeConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the details of the ConfigMap were queried. Valid values:
	//
	// 	- **true**: The details were queried.
	//
	// 	- **false**: The details failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeConfigMapResponseBody) GetData() *DescribeConfigMapResponseBodyData {
	return s.Data
}

func (s *DescribeConfigMapResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigMapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeConfigMapResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeConfigMapResponseBody) SetCode(v string) *DescribeConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetData(v *DescribeConfigMapResponseBodyData) *DescribeConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *DescribeConfigMapResponseBody) SetErrorCode(v string) *DescribeConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetMessage(v string) *DescribeConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetRequestId(v string) *DescribeConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetSuccess(v bool) *DescribeConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetTraceId(v string) *DescribeConfigMapResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeConfigMapResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeConfigMapResponseBodyData struct {
	// The ID of the ConfigMap.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The time when the ConfigMap was created.
	//
	// example:
	//
	// 1593746835111
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The key-value pairs of the ConfigMap. Format:
	//
	// {"k1":"v1", "k2":"v2"}
	//
	// k specifies a key and v specifies a value. For more information, see [Manage a Kubernetes ConfigMap](https://help.aliyun.com/document_detail/171326.html).
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description of the ConfigMap.
	//
	// example:
	//
	// test-desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the ConfigMap.
	//
	// example:
	//
	// test-configmap
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The application that is associated with the ConfigMap.
	RelateApps []*DescribeConfigMapResponseBodyDataRelateApps `json:"RelateApps,omitempty" xml:"RelateApps,omitempty" type:"Repeated"`
	// The time when the ConfigMap was updated.
	//
	// example:
	//
	// 1593747274195
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeConfigMapResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponseBodyData) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DescribeConfigMapResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeConfigMapResponseBodyData) GetData() map[string]interface{} {
	return s.Data
}

func (s *DescribeConfigMapResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeConfigMapResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeConfigMapResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeConfigMapResponseBodyData) GetRelateApps() []*DescribeConfigMapResponseBodyDataRelateApps {
	return s.RelateApps
}

func (s *DescribeConfigMapResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeConfigMapResponseBodyData) SetConfigMapId(v int64) *DescribeConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetCreateTime(v int64) *DescribeConfigMapResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetData(v map[string]interface{}) *DescribeConfigMapResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetDescription(v string) *DescribeConfigMapResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetName(v string) *DescribeConfigMapResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetNamespaceId(v string) *DescribeConfigMapResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetRelateApps(v []*DescribeConfigMapResponseBodyDataRelateApps) *DescribeConfigMapResponseBodyData {
	s.RelateApps = v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetUpdateTime(v int64) *DescribeConfigMapResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) Validate() error {
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

type DescribeConfigMapResponseBodyDataRelateApps struct {
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

func (s DescribeConfigMapResponseBodyDataRelateApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigMapResponseBodyDataRelateApps) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) GetAppId() *string {
	return s.AppId
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) GetAppName() *string {
	return s.AppName
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) SetAppId(v string) *DescribeConfigMapResponseBodyDataRelateApps {
	s.AppId = &v
	return s
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) SetAppName(v string) *DescribeConfigMapResponseBodyDataRelateApps {
	s.AppName = &v
	return s
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) Validate() error {
	return dara.Validate(s)
}
