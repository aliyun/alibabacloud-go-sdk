// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RemoveKnowledgeTagsResponseBodyData) *RemoveKnowledgeTagsResponseBody
	GetData() *RemoveKnowledgeTagsResponseBodyData
	SetRequestId(v string) *RemoveKnowledgeTagsResponseBody
	GetRequestId() *string
}

type RemoveKnowledgeTagsResponseBody struct {
	// The returned data.
	Data *RemoveKnowledgeTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveKnowledgeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeTagsResponseBody) GetData() *RemoveKnowledgeTagsResponseBodyData {
	return s.Data
}

func (s *RemoveKnowledgeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveKnowledgeTagsResponseBody) SetData(v *RemoveKnowledgeTagsResponseBodyData) *RemoveKnowledgeTagsResponseBody {
	s.Data = v
	return s
}

func (s *RemoveKnowledgeTagsResponseBody) SetRequestId(v string) *RemoveKnowledgeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveKnowledgeTagsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveKnowledgeTagsResponseBodyData struct {
	// The location of the knowledge base file.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The message returned by the operation.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of tags that were successfully deleted.
	//
	// example:
	//
	// 1
	Removed *int32 `json:"Removed,omitempty" xml:"Removed,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveKnowledgeTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeTagsResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *RemoveKnowledgeTagsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RemoveKnowledgeTagsResponseBodyData) GetRemoved() *int32 {
	return s.Removed
}

func (s *RemoveKnowledgeTagsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveKnowledgeTagsResponseBodyData) SetFileLocation(v string) *RemoveKnowledgeTagsResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *RemoveKnowledgeTagsResponseBodyData) SetMessage(v string) *RemoveKnowledgeTagsResponseBodyData {
	s.Message = &v
	return s
}

func (s *RemoveKnowledgeTagsResponseBodyData) SetRemoved(v int32) *RemoveKnowledgeTagsResponseBodyData {
	s.Removed = &v
	return s
}

func (s *RemoveKnowledgeTagsResponseBodyData) SetSuccess(v bool) *RemoveKnowledgeTagsResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveKnowledgeTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
