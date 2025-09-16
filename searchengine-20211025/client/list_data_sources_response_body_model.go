// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataSourcesResponseBodyResult) *ListDataSourcesResponseBody
	GetResult() []*ListDataSourcesResponseBodyResult
}

type ListDataSourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result []*ListDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) GetResult() []*ListDataSourcesResponseBodyResult {
	return s.Result
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetResult(v []*ListDataSourcesResponseBodyResult) *ListDataSourcesResponseBody {
	s.Result = v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyResult struct {
	// The data center in which the data source is deployed.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The time when the full data of the data source was last queried.
	//
	// example:
	//
	// 1718787785
	LastFulTime *int64 `json:"lastFulTime,omitempty" xml:"lastFulTime,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// data_source_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data source. Valid values: new: The data source is being created. publish: The data source is in the normal state. trash: The data source is being deleted.
	//
	// example:
	//
	// new
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataSourcesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ListDataSourcesResponseBodyResult) GetIndexes() []*string {
	return s.Indexes
}

func (s *ListDataSourcesResponseBodyResult) GetLastFulTime() *int64 {
	return s.LastFulTime
}

func (s *ListDataSourcesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDataSourcesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListDataSourcesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDataSourcesResponseBodyResult) SetDomain(v string) *ListDataSourcesResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetIndexes(v []*string) *ListDataSourcesResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetLastFulTime(v int64) *ListDataSourcesResponseBodyResult {
	s.LastFulTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetName(v string) *ListDataSourcesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetStatus(v string) *ListDataSourcesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) SetType(v string) *ListDataSourcesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataSourcesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
