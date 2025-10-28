// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sConfigMapsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListK8sConfigMapsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListK8sConfigMapsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8sConfigMapsResponseBody
	GetRequestId() *string
	SetResult(v *ListK8sConfigMapsResponseBodyResult) *ListK8sConfigMapsResponseBody
	GetResult() *ListK8sConfigMapsResponseBodyResult
}

type ListK8sConfigMapsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query results that are returned.
	Result *ListK8sConfigMapsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListK8sConfigMapsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListK8sConfigMapsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8sConfigMapsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sConfigMapsResponseBody) GetResult() *ListK8sConfigMapsResponseBodyResult {
	return s.Result
}

func (s *ListK8sConfigMapsResponseBody) SetCode(v int32) *ListK8sConfigMapsResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8sConfigMapsResponseBody) SetMessage(v string) *ListK8sConfigMapsResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8sConfigMapsResponseBody) SetRequestId(v string) *ListK8sConfigMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sConfigMapsResponseBody) SetResult(v *ListK8sConfigMapsResponseBodyResult) *ListK8sConfigMapsResponseBody {
	s.Result = v
	return s
}

func (s *ListK8sConfigMapsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListK8sConfigMapsResponseBodyResult struct {
	// The information about ConfigMaps.
	ConfigMaps []*ListK8sConfigMapsResponseBodyResultConfigMaps `json:"ConfigMaps,omitempty" xml:"ConfigMaps,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 6
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListK8sConfigMapsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsResponseBodyResult) GetConfigMaps() []*ListK8sConfigMapsResponseBodyResultConfigMaps {
	return s.ConfigMaps
}

func (s *ListK8sConfigMapsResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListK8sConfigMapsResponseBodyResult) SetConfigMaps(v []*ListK8sConfigMapsResponseBodyResultConfigMaps) *ListK8sConfigMapsResponseBodyResult {
	s.ConfigMaps = v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResult) SetTotal(v int32) *ListK8sConfigMapsResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResult) Validate() error {
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

type ListK8sConfigMapsResponseBodyResultConfigMaps struct {
	// The ID of the Kubernetes cluster. You can obtain the cluster ID by calling the GetK8sCluster operation. For more information, see [GetK8sCluster](https://help.aliyun.com/document_detail/181437.html).
	//
	// example:
	//
	// d73918f4-3b08-4c17-bb07-eaf8********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// my-cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The time when the ConfigMaps were created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-31T02:46:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The information about ConfigMaps.
	Data []*ListK8sConfigMapsResponseBodyResultConfigMapsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The name of the ConfigMap.
	//
	// example:
	//
	// my-config
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The related applications.
	RelatedApps []*ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps `json:"RelatedApps,omitempty" xml:"RelatedApps,omitempty" type:"Repeated"`
}

func (s ListK8sConfigMapsResponseBodyResultConfigMaps) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsResponseBodyResultConfigMaps) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetData() []*ListK8sConfigMapsResponseBodyResultConfigMapsData {
	return s.Data
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetName() *string {
	return s.Name
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) GetRelatedApps() []*ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps {
	return s.RelatedApps
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetClusterId(v string) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.ClusterId = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetClusterName(v string) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.ClusterName = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetCreationTime(v string) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.CreationTime = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetData(v []*ListK8sConfigMapsResponseBodyResultConfigMapsData) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.Data = v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetName(v string) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.Name = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetNamespace(v string) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.Namespace = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) SetRelatedApps(v []*ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) *ListK8sConfigMapsResponseBodyResultConfigMaps {
	s.RelatedApps = v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMaps) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedApps != nil {
		for _, item := range s.RelatedApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sConfigMapsResponseBodyResultConfigMapsData struct {
	// The user-defined key that is stored in the ConfigMap.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The user-defined value that is stored in the ConfigMap.
	//
	// example:
	//
	// william
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListK8sConfigMapsResponseBodyResultConfigMapsData) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsResponseBodyResultConfigMapsData) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsData) GetKey() *string {
	return s.Key
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsData) GetValue() *string {
	return s.Value
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsData) SetKey(v string) *ListK8sConfigMapsResponseBodyResultConfigMapsData {
	s.Key = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsData) SetValue(v string) *ListK8sConfigMapsResponseBodyResultConfigMapsData {
	s.Value = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsData) Validate() error {
	return dara.Validate(s)
}

type ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps struct {
	// The ID of the application.
	//
	// example:
	//
	// 728cbdf2-da10-49b8-b69c-9168a********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// my-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) GetAppId() *string {
	return s.AppId
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) GetAppName() *string {
	return s.AppName
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) SetAppId(v string) *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps {
	s.AppId = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) SetAppName(v string) *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps {
	s.AppName = &v
	return s
}

func (s *ListK8sConfigMapsResponseBodyResultConfigMapsRelatedApps) Validate() error {
	return dara.Validate(s)
}
