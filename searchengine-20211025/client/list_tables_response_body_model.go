// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetResult(v []*ListTablesResponseBodyResult) *ListTablesResponseBody
	GetResult() []*ListTablesResponseBodyResult
}

type ListTablesResponseBody struct {
	// requestId
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result []*ListTablesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetResult() []*ListTablesResponseBodyResult {
	return s.Result
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetResult(v []*ListTablesResponseBodyResult) *ListTablesResponseBody {
	s.Result = v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyResult struct {
	// The state of the index table. Valid values: NEW, PUBLISH, IN_USE, NOT_USE, STOP_USE, RESTORE_USE, and FAIL. After an index is created in an OpenSearch Retrieval Engine Edition instance, the index enters the IN_USE state. If the first full index fails to be created in an OpenSearch Vector Search Edition instance of the new version, the index is in the FAIL state.
	//
	// example:
	//
	// IN_USE
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	// The index name.
	//
	// example:
	//
	// es_test_1b
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The state of the index table. Valid values: NEW, PUBLISH, IN_USE, NOT_USE, STOP_USE, RESTORE_USE, and FAIL. After an index is created in an OpenSearch Retrieval Engine Edition instance, the index enters the IN_USE state. If the first full index fails to be created in an OpenSearch Vector Search Edition instance of the new version, the index is in the FAIL state.
	//
	// example:
	//
	// IN_USE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListTablesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyResult) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *ListTablesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListTablesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListTablesResponseBodyResult) SetIndexStatus(v string) *ListTablesResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *ListTablesResponseBodyResult) SetName(v string) *ListTablesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyResult) SetStatus(v string) *ListTablesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListTablesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
