// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataStreamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListDataStreamsResponseBodyHeaders) *ListDataStreamsResponseBody
	GetHeaders() *ListDataStreamsResponseBodyHeaders
	SetRequestId(v string) *ListDataStreamsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataStreamsResponseBodyResult) *ListDataStreamsResponseBody
	GetResult() []*ListDataStreamsResponseBodyResult
}

type ListDataStreamsResponseBody struct {
	Headers *ListDataStreamsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDataStreamsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDataStreamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBody) GetHeaders() *ListDataStreamsResponseBodyHeaders {
	return s.Headers
}

func (s *ListDataStreamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataStreamsResponseBody) GetResult() []*ListDataStreamsResponseBodyResult {
	return s.Result
}

func (s *ListDataStreamsResponseBody) SetHeaders(v *ListDataStreamsResponseBodyHeaders) *ListDataStreamsResponseBody {
	s.Headers = v
	return s
}

func (s *ListDataStreamsResponseBody) SetRequestId(v string) *ListDataStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataStreamsResponseBody) SetResult(v []*ListDataStreamsResponseBodyResult) *ListDataStreamsResponseBody {
	s.Result = v
	return s
}

func (s *ListDataStreamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataStreamsResponseBodyHeaders struct {
	// example:
	//
	// 100
	XManagedCount *int32 `json:"X-Managed-Count,omitempty" xml:"X-Managed-Count,omitempty"`
	// example:
	//
	// 143993923932990
	XManagedStorageSize *int64 `json:"X-Managed-StorageSize,omitempty" xml:"X-Managed-StorageSize,omitempty"`
}

func (s ListDataStreamsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDataStreamsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyHeaders) GetXManagedCount() *int32 {
	return s.XManagedCount
}

func (s *ListDataStreamsResponseBodyHeaders) GetXManagedStorageSize() *int64 {
	return s.XManagedStorageSize
}

func (s *ListDataStreamsResponseBodyHeaders) SetXManagedCount(v int32) *ListDataStreamsResponseBodyHeaders {
	s.XManagedCount = &v
	return s
}

func (s *ListDataStreamsResponseBodyHeaders) SetXManagedStorageSize(v int64) *ListDataStreamsResponseBodyHeaders {
	s.XManagedStorageSize = &v
	return s
}

func (s *ListDataStreamsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListDataStreamsResponseBodyResult struct {
	// example:
	//
	// Green
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// example:
	//
	// rollver1
	IlmPolicyName *string `json:"ilmPolicyName,omitempty" xml:"ilmPolicyName,omitempty"`
	// example:
	//
	// template1
	IndexTemplateName *string                                     `json:"indexTemplateName,omitempty" xml:"indexTemplateName,omitempty"`
	Indices           []*ListDataStreamsResponseBodyResultIndices `json:"indices,omitempty" xml:"indices,omitempty" type:"Repeated"`
	// example:
	//
	// 1788239393298
	ManagedStorageSize *int64 `json:"managedStorageSize,omitempty" xml:"managedStorageSize,omitempty"`
	// example:
	//
	// my-index-0001
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1788239393298
	TotalStorageSize *int64 `json:"totalStorageSize,omitempty" xml:"totalStorageSize,omitempty"`
}

func (s ListDataStreamsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataStreamsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyResult) GetHealth() *string {
	return s.Health
}

func (s *ListDataStreamsResponseBodyResult) GetIlmPolicyName() *string {
	return s.IlmPolicyName
}

func (s *ListDataStreamsResponseBodyResult) GetIndexTemplateName() *string {
	return s.IndexTemplateName
}

func (s *ListDataStreamsResponseBodyResult) GetIndices() []*ListDataStreamsResponseBodyResultIndices {
	return s.Indices
}

func (s *ListDataStreamsResponseBodyResult) GetManagedStorageSize() *int64 {
	return s.ManagedStorageSize
}

func (s *ListDataStreamsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDataStreamsResponseBodyResult) GetTotalStorageSize() *int64 {
	return s.TotalStorageSize
}

func (s *ListDataStreamsResponseBodyResult) SetHealth(v string) *ListDataStreamsResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetIlmPolicyName(v string) *ListDataStreamsResponseBodyResult {
	s.IlmPolicyName = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetIndexTemplateName(v string) *ListDataStreamsResponseBodyResult {
	s.IndexTemplateName = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetIndices(v []*ListDataStreamsResponseBodyResultIndices) *ListDataStreamsResponseBodyResult {
	s.Indices = v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetManagedStorageSize(v int64) *ListDataStreamsResponseBodyResult {
	s.ManagedStorageSize = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetName(v string) *ListDataStreamsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetTotalStorageSize(v int64) *ListDataStreamsResponseBodyResult {
	s.TotalStorageSize = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDataStreamsResponseBodyResultIndices struct {
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// Green
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// example:
	//
	// false
	IsManaged *bool `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	// example:
	//
	// following
	ManagedStatus *string `json:"managedStatus,omitempty" xml:"managedStatus,omitempty"`
	// example:
	//
	// Log1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 15393899
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListDataStreamsResponseBodyResultIndices) String() string {
	return dara.Prettify(s)
}

func (s ListDataStreamsResponseBodyResultIndices) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyResultIndices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataStreamsResponseBodyResultIndices) GetHealth() *string {
	return s.Health
}

func (s *ListDataStreamsResponseBodyResultIndices) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListDataStreamsResponseBodyResultIndices) GetManagedStatus() *string {
	return s.ManagedStatus
}

func (s *ListDataStreamsResponseBodyResultIndices) GetName() *string {
	return s.Name
}

func (s *ListDataStreamsResponseBodyResultIndices) GetSize() *int64 {
	return s.Size
}

func (s *ListDataStreamsResponseBodyResultIndices) SetCreateTime(v string) *ListDataStreamsResponseBodyResultIndices {
	s.CreateTime = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetHealth(v string) *ListDataStreamsResponseBodyResultIndices {
	s.Health = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetIsManaged(v bool) *ListDataStreamsResponseBodyResultIndices {
	s.IsManaged = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetManagedStatus(v string) *ListDataStreamsResponseBodyResultIndices {
	s.ManagedStatus = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetName(v string) *ListDataStreamsResponseBodyResultIndices {
	s.Name = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetSize(v int64) *ListDataStreamsResponseBodyResultIndices {
	s.Size = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) Validate() error {
	return dara.Validate(s)
}
