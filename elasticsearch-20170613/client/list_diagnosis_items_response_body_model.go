// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDiagnosisItemsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDiagnosisItemsResponseBodyResult) *ListDiagnosisItemsResponseBody
	GetResult() []*ListDiagnosisItemsResponseBodyResult
}

type ListDiagnosisItemsResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDiagnosisItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnosisItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosisItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnosisItemsResponseBody) GetResult() []*ListDiagnosisItemsResponseBodyResult {
	return s.Result
}

func (s *ListDiagnosisItemsResponseBody) SetRequestId(v string) *ListDiagnosisItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosisItemsResponseBody) SetResult(v []*ListDiagnosisItemsResponseBodyResult) *ListDiagnosisItemsResponseBody {
	s.Result = v
	return s
}

func (s *ListDiagnosisItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDiagnosisItemsResponseBodyResult struct {
	// example:
	//
	// 诊断集群写数据是否有堆积当集群的数据写入存在堆积时，会造成BulkReject异常，可能会导致数据丢失，且会造成系统资源消耗严重
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ClusterBulkRejectDiagnostic
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 索引写入BulkReject诊断
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListDiagnosisItemsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDiagnosisItemsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListDiagnosisItemsResponseBodyResult) GetKey() *string {
	return s.Key
}

func (s *ListDiagnosisItemsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDiagnosisItemsResponseBodyResult) SetDescription(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetKey(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Key = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) SetName(v string) *ListDiagnosisItemsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDiagnosisItemsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
