// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDataSourceResponseBody
	GetRequestId() *string
	SetResult(v *GetDataSourceResponseBodyResult) *GetDataSourceResponseBody
	GetResult() *GetDataSourceResponseBodyResult
}

type GetDataSourceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the data source.
	Result *GetDataSourceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceResponseBody) GetResult() *GetDataSourceResponseBodyResult {
	return s.Result
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetResult(v *GetDataSourceResponseBodyResult) *GetDataSourceResponseBody {
	s.Result = v
	return s
}

func (s *GetDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceResponseBodyResult struct {
	// The data center where the data source is deployed in offline mode
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The list of index information
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The time when the full data of the data source was last queried.
	//
	// example:
	//
	// 1718787219
	LastFulTime *int64 `json:"lastFulTime,omitempty" xml:"lastFulTime,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// my_index
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data source. Valid values: new: The data source is being created. publish: The data source is in the normal state. trash: The data source is being deleted.
	//
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the data source
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDataSourceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *GetDataSourceResponseBodyResult) GetIndexes() []*string {
	return s.Indexes
}

func (s *GetDataSourceResponseBodyResult) GetLastFulTime() *int64 {
	return s.LastFulTime
}

func (s *GetDataSourceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetDataSourceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetDataSourceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetDataSourceResponseBodyResult) SetDomain(v string) *GetDataSourceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetIndexes(v []*string) *GetDataSourceResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetLastFulTime(v int64) *GetDataSourceResponseBodyResult {
	s.LastFulTime = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetName(v string) *GetDataSourceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetStatus(v string) *GetDataSourceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) SetType(v string) *GetDataSourceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetDataSourceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
