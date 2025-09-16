// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetFileResponseBody
	GetRequestId() *string
	SetResult(v *GetFileResponseBodyResult) *GetFileResponseBody
	GetResult() *GetFileResponseBodyResult
}

type GetFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index information.
	Result *GetFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileResponseBody) GetResult() *GetFileResponseBodyResult {
	return s.Result
}

func (s *GetFileResponseBody) SetRequestId(v string) *GetFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileResponseBody) SetResult(v *GetFileResponseBodyResult) *GetFileResponseBody {
	s.Result = v
	return s
}

func (s *GetFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileResponseBodyResult struct {
	// The file content.
	//
	// example:
	//
	// None
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The data source.
	//
	// example:
	//
	// ha-cn-2r42p5oi202_xijie_test
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// Extended information
	Extend map[string][]*string `json:"extend,omitempty" xml:"extend,omitempty"`
	// The full path of the file.
	//
	// example:
	//
	// ""
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// The file name.
	//
	// example:
	//
	// ha-cn-2r42ostoc01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// ds=20210828
	Partition *int64 `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetFileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetFileResponseBodyResult) GetDataSource() *string {
	return s.DataSource
}

func (s *GetFileResponseBodyResult) GetExtend() map[string][]*string {
	return s.Extend
}

func (s *GetFileResponseBodyResult) GetFullPathName() *string {
	return s.FullPathName
}

func (s *GetFileResponseBodyResult) GetIsDir() *bool {
	return s.IsDir
}

func (s *GetFileResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetFileResponseBodyResult) GetPartition() *int64 {
	return s.Partition
}

func (s *GetFileResponseBodyResult) SetContent(v string) *GetFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFileResponseBodyResult) SetDataSource(v string) *GetFileResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *GetFileResponseBodyResult) SetExtend(v map[string][]*string) *GetFileResponseBodyResult {
	s.Extend = v
	return s
}

func (s *GetFileResponseBodyResult) SetFullPathName(v string) *GetFileResponseBodyResult {
	s.FullPathName = &v
	return s
}

func (s *GetFileResponseBodyResult) SetIsDir(v bool) *GetFileResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *GetFileResponseBodyResult) SetName(v string) *GetFileResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetFileResponseBodyResult) SetPartition(v int64) *GetFileResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *GetFileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
