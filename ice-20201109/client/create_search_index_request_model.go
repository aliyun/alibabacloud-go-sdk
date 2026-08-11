// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexConfig(v string) *CreateSearchIndexRequest
	GetIndexConfig() *string
	SetIndexStatus(v string) *CreateSearchIndexRequest
	GetIndexStatus() *string
	SetIndexType(v string) *CreateSearchIndexRequest
	GetIndexType() *string
	SetSearchLibName(v string) *CreateSearchIndexRequest
	GetSearchLibName() *string
}

type CreateSearchIndexRequest struct {
	// The index configuration.
	//
	// example:
	//
	// {}
	IndexConfig *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// The index status. Default value: Active. Valid values:
	//
	// - Active: activated.
	//
	// - Deactive: deactivated.
	//
	// example:
	//
	// Active
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// The index type. Valid values:
	//
	// - mm: large model visual state. Used to describe complex visual features and actions in videos. This type helps identify and search for specific actions, movements, and events in videos, such as a soccer player scoring a goal or a basketball player getting injured.
	//
	// >
	//
	// > The shared instance type supports up to 1,000 hours of video. After the limit is exceeded, the system no longer performs large model visual state analysis.
	//
	// - face: automatic face recognition. Used to describe facial features in videos. Through face recognition technology, faces in videos can be automatically tagged and searched.
	//
	// >
	//
	// > The shared instance type supports up to 1,000,000 face analyses. After the limit is exceeded, the system no longer performs face analysis.
	//
	// - aiLabel: intelligent tagging. The intelligent tagging index type is used to describe subtitles, speech, and other content in videos. Through text and speech recognition technology, language information such as subtitles and dialogues in videos can be automatically extracted for tagging and searching. This helps users quickly search for and locate content related to specific topics or keywords in videos.
	//
	// This parameter is required.
	//
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The name of the search library.
	//
	// - If you do not specify a search library name, the search index is created in the default IMS search library. Default value: ims-default-search-lib.
	//
	// - You can call the [QuerySearchLib](https://help.aliyun.com/document_detail/2584455.html) operation to query information about existing search libraries.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s CreateSearchIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchIndexRequest) GetIndexConfig() *string {
	return s.IndexConfig
}

func (s *CreateSearchIndexRequest) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *CreateSearchIndexRequest) GetIndexType() *string {
	return s.IndexType
}

func (s *CreateSearchIndexRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *CreateSearchIndexRequest) SetIndexConfig(v string) *CreateSearchIndexRequest {
	s.IndexConfig = &v
	return s
}

func (s *CreateSearchIndexRequest) SetIndexStatus(v string) *CreateSearchIndexRequest {
	s.IndexStatus = &v
	return s
}

func (s *CreateSearchIndexRequest) SetIndexType(v string) *CreateSearchIndexRequest {
	s.IndexType = &v
	return s
}

func (s *CreateSearchIndexRequest) SetSearchLibName(v string) *CreateSearchIndexRequest {
	s.SearchLibName = &v
	return s
}

func (s *CreateSearchIndexRequest) Validate() error {
	return dara.Validate(s)
}
