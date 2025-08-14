// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchIndexJobRerunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *SearchIndexJobRerunRequest
	GetMediaIds() *string
	SetNamespace(v string) *SearchIndexJobRerunRequest
	GetNamespace() *string
	SetSearchLibName(v string) *SearchIndexJobRerunRequest
	GetSearchLibName() *string
	SetTask(v string) *SearchIndexJobRerunRequest
	GetTask() *string
}

type SearchIndexJobRerunRequest struct {
	// The ID of the media asset. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******c48fb37407365d4f2cd8******
	MediaIds  *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The search library.
	//
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The type of the job. Separate multiple types with commas (,).
	//
	// 	- aiLabel: smart tagging.
	//
	// 	- face: face recognition.
	//
	// 	- mm: large visual model.
	//
	// example:
	//
	// AiLabel,Face,Mm
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s SearchIndexJobRerunRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchIndexJobRerunRequest) GoString() string {
	return s.String()
}

func (s *SearchIndexJobRerunRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *SearchIndexJobRerunRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchIndexJobRerunRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchIndexJobRerunRequest) GetTask() *string {
	return s.Task
}

func (s *SearchIndexJobRerunRequest) SetMediaIds(v string) *SearchIndexJobRerunRequest {
	s.MediaIds = &v
	return s
}

func (s *SearchIndexJobRerunRequest) SetNamespace(v string) *SearchIndexJobRerunRequest {
	s.Namespace = &v
	return s
}

func (s *SearchIndexJobRerunRequest) SetSearchLibName(v string) *SearchIndexJobRerunRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchIndexJobRerunRequest) SetTask(v string) *SearchIndexJobRerunRequest {
	s.Task = &v
	return s
}

func (s *SearchIndexJobRerunRequest) Validate() error {
	return dara.Validate(s)
}
