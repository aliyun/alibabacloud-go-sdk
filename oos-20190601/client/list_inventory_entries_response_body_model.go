// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventoryEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaptureTime(v string) *ListInventoryEntriesResponseBody
	GetCaptureTime() *string
	SetEntries(v []map[string]interface{}) *ListInventoryEntriesResponseBody
	GetEntries() []map[string]interface{}
	SetInstanceId(v string) *ListInventoryEntriesResponseBody
	GetInstanceId() *string
	SetMaxResults(v int32) *ListInventoryEntriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInventoryEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInventoryEntriesResponseBody
	GetRequestId() *string
	SetSchemaVersion(v string) *ListInventoryEntriesResponseBody
	GetSchemaVersion() *string
	SetTypeName(v string) *ListInventoryEntriesResponseBody
	GetTypeName() *string
}

type ListInventoryEntriesResponseBody struct {
	// The time when the request was sent.
	//
	// example:
	//
	// 2020-09-17T12:28:13Z
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The configurations of the component.
	Entries []map[string]interface{} `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp1cpoxxxwxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A81E4B2E-6B33-4BAE-9856-55DB7C893E01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version number of the component.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// ACS:InstanceInformation
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListInventoryEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesResponseBody) GetCaptureTime() *string {
	return s.CaptureTime
}

func (s *ListInventoryEntriesResponseBody) GetEntries() []map[string]interface{} {
	return s.Entries
}

func (s *ListInventoryEntriesResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInventoryEntriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInventoryEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInventoryEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInventoryEntriesResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *ListInventoryEntriesResponseBody) GetTypeName() *string {
	return s.TypeName
}

func (s *ListInventoryEntriesResponseBody) SetCaptureTime(v string) *ListInventoryEntriesResponseBody {
	s.CaptureTime = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetEntries(v []map[string]interface{}) *ListInventoryEntriesResponseBody {
	s.Entries = v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetInstanceId(v string) *ListInventoryEntriesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetMaxResults(v int32) *ListInventoryEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetNextToken(v string) *ListInventoryEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetRequestId(v string) *ListInventoryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetSchemaVersion(v string) *ListInventoryEntriesResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetTypeName(v string) *ListInventoryEntriesResponseBody {
	s.TypeName = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
