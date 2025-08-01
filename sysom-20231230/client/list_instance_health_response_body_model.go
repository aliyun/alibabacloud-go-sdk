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
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListInstanceHealthResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// Query no data
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
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
	return dara.Validate(s)
}

type ListInstanceHealthResponseBodyData struct {
	Images    []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Instance  *string   `json:"instance,omitempty" xml:"instance,omitempty"`
	Namespace *string   `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Pod       *string   `json:"pod,omitempty" xml:"pod,omitempty"`
	RegionId  *string   `json:"region_id,omitempty" xml:"region_id,omitempty"`
	Score     *float32  `json:"score,omitempty" xml:"score,omitempty"`
	Status    *string   `json:"status,omitempty" xml:"status,omitempty"`
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
