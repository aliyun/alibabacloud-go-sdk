// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDataInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetClusterDataInformationResponseBody
	GetRequestId() *string
	SetResult(v *GetClusterDataInformationResponseBodyResult) *GetClusterDataInformationResponseBody
	GetResult() *GetClusterDataInformationResponseBodyResult
}

type GetClusterDataInformationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result *GetClusterDataInformationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetClusterDataInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDataInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterDataInformationResponseBody) GetResult() *GetClusterDataInformationResponseBodyResult {
	return s.Result
}

func (s *GetClusterDataInformationResponseBody) SetRequestId(v string) *GetClusterDataInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterDataInformationResponseBody) SetResult(v *GetClusterDataInformationResponseBodyResult) *GetClusterDataInformationResponseBody {
	s.Result = v
	return s
}

func (s *GetClusterDataInformationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterDataInformationResponseBodyResult struct {
	// Whether it is connectable.
	//
	// example:
	//
	// true
	Connectable *bool `json:"connectable,omitempty" xml:"connectable,omitempty"`
	// The metadata of the cluster.
	MetaInfo *GetClusterDataInformationResponseBodyResultMetaInfo `json:"metaInfo,omitempty" xml:"metaInfo,omitempty" type:"Struct"`
}

func (s GetClusterDataInformationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDataInformationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponseBodyResult) GetConnectable() *bool {
	return s.Connectable
}

func (s *GetClusterDataInformationResponseBodyResult) GetMetaInfo() *GetClusterDataInformationResponseBodyResultMetaInfo {
	return s.MetaInfo
}

func (s *GetClusterDataInformationResponseBodyResult) SetConnectable(v bool) *GetClusterDataInformationResponseBodyResult {
	s.Connectable = &v
	return s
}

func (s *GetClusterDataInformationResponseBodyResult) SetMetaInfo(v *GetClusterDataInformationResponseBodyResultMetaInfo) *GetClusterDataInformationResponseBodyResult {
	s.MetaInfo = v
	return s
}

func (s *GetClusterDataInformationResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetClusterDataInformationResponseBodyResultMetaInfo struct {
	// The fields in the Mapping for the index.
	Fields []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The index list of the cluster.
	Indices []*string `json:"indices,omitempty" xml:"indices,omitempty" type:"Repeated"`
	// The Mapping configuration of the cluster.
	//
	// example:
	//
	// {\\"_doc\\":{\\"properties\\":{\\"user\\":{\\"properties\\":{\\"last\\":{\\"type\\":\\"text\\",...}}}}}}
	Mapping *string `json:"mapping,omitempty" xml:"mapping,omitempty"`
	// The Settings of the cluster.
	//
	// example:
	//
	// {\\n  \\"index\\": {\\n    \\"replication\\": {\\n}.....}}
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
	// Specifies the type of the index.
	TypeName []*string `json:"typeName,omitempty" xml:"typeName,omitempty" type:"Repeated"`
}

func (s GetClusterDataInformationResponseBodyResultMetaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDataInformationResponseBodyResultMetaInfo) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) GetFields() []*string {
	return s.Fields
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) GetIndices() []*string {
	return s.Indices
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) GetMapping() *string {
	return s.Mapping
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) GetSettings() *string {
	return s.Settings
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) GetTypeName() []*string {
	return s.TypeName
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetFields(v []*string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Fields = v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetIndices(v []*string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Indices = v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetMapping(v string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Mapping = &v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetSettings(v string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Settings = &v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetTypeName(v []*string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.TypeName = v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) Validate() error {
	return dara.Validate(s)
}
