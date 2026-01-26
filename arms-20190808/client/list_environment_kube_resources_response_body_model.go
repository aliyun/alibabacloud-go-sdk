// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentKubeResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentKubeResourcesResponseBody
	GetCode() *int32
	SetData(v []*ListEnvironmentKubeResourcesResponseBodyData) *ListEnvironmentKubeResourcesResponseBody
	GetData() []*ListEnvironmentKubeResourcesResponseBodyData
	SetMessage(v string) *ListEnvironmentKubeResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentKubeResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentKubeResourcesResponseBody
	GetSuccess() *bool
}

type ListEnvironmentKubeResourcesResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data []*ListEnvironmentKubeResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentKubeResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentKubeResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentKubeResourcesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentKubeResourcesResponseBody) GetData() []*ListEnvironmentKubeResourcesResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentKubeResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentKubeResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentKubeResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentKubeResourcesResponseBody) SetCode(v int32) *ListEnvironmentKubeResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBody) SetData(v []*ListEnvironmentKubeResourcesResponseBodyData) *ListEnvironmentKubeResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBody) SetMessage(v string) *ListEnvironmentKubeResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBody) SetRequestId(v string) *ListEnvironmentKubeResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBody) SetSuccess(v bool) *ListEnvironmentKubeResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvironmentKubeResourcesResponseBodyData struct {
	// The version number of the API.
	//
	// example:
	//
	// v1
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Pod
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The metadata.
	Metadata *ListEnvironmentKubeResourcesResponseBodyDataMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// The resource specifications.
	//
	// example:
	//
	// {
	//
	//         "dnsPolicy": "ClusterFirst",
	//
	//         "nodeName": "cn-hangzhou.172.16.0.60",
	//
	//         "terminationGracePeriodSeconds": 30,
	//
	//         "enableServiceLinks": true,
	//
	//         "serviceAccountName": "arms-prom-operator",
	//
	//         "volumes": [
	//
	//           {
	//
	//             "name": "certs",
	//
	//             "secret": {
	//
	//               "secretName": "arms-prometheus-ack-arms-prometheus-cert",
	//
	//               "defaultMode": 420
	//
	//             }
	//
	//           }
	Spec interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// run
	Status interface{} `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEnvironmentKubeResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentKubeResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) GetKind() *string {
	return s.Kind
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) GetMetadata() *ListEnvironmentKubeResourcesResponseBodyDataMetadata {
	return s.Metadata
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) GetSpec() interface{} {
	return s.Spec
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) GetStatus() interface{} {
	return s.Status
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) SetApiVersion(v string) *ListEnvironmentKubeResourcesResponseBodyData {
	s.ApiVersion = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) SetKind(v string) *ListEnvironmentKubeResourcesResponseBodyData {
	s.Kind = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) SetMetadata(v *ListEnvironmentKubeResourcesResponseBodyDataMetadata) *ListEnvironmentKubeResourcesResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) SetSpec(v interface{}) *ListEnvironmentKubeResourcesResponseBodyData {
	s.Spec = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) SetStatus(v interface{}) *ListEnvironmentKubeResourcesResponseBodyData {
	s.Status = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyData) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEnvironmentKubeResourcesResponseBodyDataMetadata struct {
	// The annotations.
	Annotations map[string]*string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The tags.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The resource name.
	//
	// example:
	//
	// arms-prometheus-ack-arms-prometheus-c577b6cc8-mvdwd
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListEnvironmentKubeResourcesResponseBodyDataMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentKubeResourcesResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) SetAnnotations(v map[string]*string) *ListEnvironmentKubeResourcesResponseBodyDataMetadata {
	s.Annotations = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) SetLabels(v map[string]*string) *ListEnvironmentKubeResourcesResponseBodyDataMetadata {
	s.Labels = v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) SetName(v string) *ListEnvironmentKubeResourcesResponseBodyDataMetadata {
	s.Name = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) SetNamespace(v string) *ListEnvironmentKubeResourcesResponseBodyDataMetadata {
	s.Namespace = &v
	return s
}

func (s *ListEnvironmentKubeResourcesResponseBodyDataMetadata) Validate() error {
	return dara.Validate(s)
}
