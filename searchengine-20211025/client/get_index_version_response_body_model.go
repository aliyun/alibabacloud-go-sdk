// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetIndexVersionResponseBody
	GetRequestId() *string
	SetResult(v *GetIndexVersionResponseBodyResult) *GetIndexVersionResponseBody
	GetResult() *GetIndexVersionResponseBodyResult
}

type GetIndexVersionResponseBody struct {
	// id of request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The clusters.
	Result *GetIndexVersionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexVersionResponseBody) GetResult() *GetIndexVersionResponseBodyResult {
	return s.Result
}

func (s *GetIndexVersionResponseBody) SetRequestId(v string) *GetIndexVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexVersionResponseBody) SetResult(v *GetIndexVersionResponseBodyResult) *GetIndexVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetIndexVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIndexVersionResponseBodyResult struct {
	// The cluster name.
	//
	// example:
	//
	// ayoss-cn-zhangjiakou-b
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The index versions.
	IndexVersions []*GetIndexVersionResponseBodyResultIndexVersions `json:"indexVersions,omitempty" xml:"indexVersions,omitempty" type:"Repeated"`
}

func (s GetIndexVersionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetIndexVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBodyResult) GetCluster() *string {
	return s.Cluster
}

func (s *GetIndexVersionResponseBodyResult) GetIndexVersions() []*GetIndexVersionResponseBodyResultIndexVersions {
	return s.IndexVersions
}

func (s *GetIndexVersionResponseBodyResult) SetCluster(v string) *GetIndexVersionResponseBodyResult {
	s.Cluster = &v
	return s
}

func (s *GetIndexVersionResponseBodyResult) SetIndexVersions(v []*GetIndexVersionResponseBodyResultIndexVersions) *GetIndexVersionResponseBodyResult {
	s.IndexVersions = v
	return s
}

func (s *GetIndexVersionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetIndexVersionResponseBodyResultIndexVersions struct {
	// The ID of the offline deployment.
	//
	// example:
	//
	// " "
	BuildDeployId *string `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The current online version number.
	//
	// example:
	//
	// 1
	CurrentVersion *int64 `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The name of the index table.
	//
	// example:
	//
	// table4
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The index versions.
	Versions []*int64 `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetIndexVersionResponseBodyResultIndexVersions) String() string {
	return dara.Prettify(s)
}

func (s GetIndexVersionResponseBodyResultIndexVersions) GoString() string {
	return s.String()
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) GetBuildDeployId() *string {
	return s.BuildDeployId
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) GetCurrentVersion() *int64 {
	return s.CurrentVersion
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) GetIndexName() *string {
	return s.IndexName
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) GetVersions() []*int64 {
	return s.Versions
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetBuildDeployId(v string) *GetIndexVersionResponseBodyResultIndexVersions {
	s.BuildDeployId = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetCurrentVersion(v int64) *GetIndexVersionResponseBodyResultIndexVersions {
	s.CurrentVersion = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetIndexName(v string) *GetIndexVersionResponseBodyResultIndexVersions {
	s.IndexName = &v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) SetVersions(v []*int64) *GetIndexVersionResponseBodyResultIndexVersions {
	s.Versions = v
	return s
}

func (s *GetIndexVersionResponseBodyResultIndexVersions) Validate() error {
	return dara.Validate(s)
}
