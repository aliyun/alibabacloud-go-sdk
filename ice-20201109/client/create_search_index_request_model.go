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
	// example:
	//
	// {}
	IndexConfig *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// example:
	//
	// Active
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// The category of the index. Valid values:
	//
	// 	- mm: large visual model. You can use this model to describe complex visual features and identify and search for specific actions, movements, and events in videos, such as when athletes score a goal or get injured.
	//
	// >  This feature is in the public preview phase. You can use this feature for free for 1,000 hours of videos.
	//
	// 	- face: face recognition. You can use the face recognition technology to describe face characteristics and automatically mark or search for faces in videos.
	//
	// 	- aiLabel: smart tagging. The smart tagging category is used to describe content such as subtitles and audio in videos. You can use the speech recognition technology to automatically extract, mark, and search for subtitles and dialog content from videos. This helps you quickly locate the video content that is related to specific topics or keywords.
	//
	// This parameter is required.
	//
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
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
