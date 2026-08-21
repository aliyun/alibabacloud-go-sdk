// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHealthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceHealthResponseBody
	GetCode() *string
	SetData(v []*ListInstanceHealthResponseBodyData) *ListInstanceHealthResponseBody
	GetData() []*ListInstanceHealthResponseBodyData
	SetMessage(v string) *ListInstanceHealthResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceHealthResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListInstanceHealthResponseBody
	GetTotal() *int32
}

type ListInstanceHealthResponseBody struct {
	// The status code.
	//
	// - `code == Success` indicates that the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListInstanceHealthResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error information of the request.
	//
	// example:
	//
	// Query no data
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 42
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstanceHealthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceHealthResponseBody) GetData() []*ListInstanceHealthResponseBodyData {
	return s.Data
}

func (s *ListInstanceHealthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceHealthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceHealthResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListInstanceHealthResponseBody) SetCode(v string) *ListInstanceHealthResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceHealthResponseBody) SetData(v []*ListInstanceHealthResponseBodyData) *ListInstanceHealthResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceHealthResponseBody) SetMessage(v string) *ListInstanceHealthResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceHealthResponseBody) SetRequestId(v string) *ListInstanceHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceHealthResponseBody) SetTotal(v int32) *ListInstanceHealthResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstanceHealthResponseBody) Validate() error {
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

type ListInstanceHealthResponseBodyData struct {
	// The list of container image names in the Pod.
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The namespace of the Pod.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Pod name.
	//
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The health score.
	//
	// example:
	//
	// 100
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// The running status of the instance. Valid values:
	//
	// - **Running**: running.
	//
	// - **Offline**: offline.
	//
	//
	//
	// > An instance in the Offline state indicates that the heartbeat between the node and the SysOM server is lost. It does not mean that the corresponding ECS instance is not running.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceHealthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHealthResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthResponseBodyData) GetImages() []*string {
	return s.Images
}

func (s *ListInstanceHealthResponseBodyData) GetInstance() *string {
	return s.Instance
}

func (s *ListInstanceHealthResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInstanceHealthResponseBodyData) GetPod() *string {
	return s.Pod
}

func (s *ListInstanceHealthResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceHealthResponseBodyData) GetScore() *float32 {
	return s.Score
}

func (s *ListInstanceHealthResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceHealthResponseBodyData) SetImages(v []*string) *ListInstanceHealthResponseBodyData {
	s.Images = v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetInstance(v string) *ListInstanceHealthResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetNamespace(v string) *ListInstanceHealthResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetPod(v string) *ListInstanceHealthResponseBodyData {
	s.Pod = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetRegionId(v string) *ListInstanceHealthResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetScore(v float32) *ListInstanceHealthResponseBodyData {
	s.Score = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetStatus(v string) *ListInstanceHealthResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) Validate() error {
	return dara.Validate(s)
}
