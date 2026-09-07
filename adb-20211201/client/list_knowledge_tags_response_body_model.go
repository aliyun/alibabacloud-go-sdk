// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListKnowledgeTagsResponseBodyData) *ListKnowledgeTagsResponseBody
	GetData() *ListKnowledgeTagsResponseBodyData
	SetRequestId(v string) *ListKnowledgeTagsResponseBody
	GetRequestId() *string
}

type ListKnowledgeTagsResponseBody struct {
	// The returned data.
	Data *ListKnowledgeTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListKnowledgeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeTagsResponseBody) GetData() *ListKnowledgeTagsResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeTagsResponseBody) SetData(v *ListKnowledgeTagsResponseBodyData) *ListKnowledgeTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeTagsResponseBody) SetRequestId(v string) *ListKnowledgeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeTagsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeTagsResponseBodyData struct {
	// The location of the knowledge base file.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// The list details.
	Tags []*ListKnowledgeTagsResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListKnowledgeTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeTagsResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *ListKnowledgeTagsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeTagsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeTagsResponseBodyData) GetTags() []*ListKnowledgeTagsResponseBodyDataTags {
	return s.Tags
}

func (s *ListKnowledgeTagsResponseBodyData) SetFileLocation(v string) *ListKnowledgeTagsResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *ListKnowledgeTagsResponseBodyData) SetMessage(v string) *ListKnowledgeTagsResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListKnowledgeTagsResponseBodyData) SetSuccess(v bool) *ListKnowledgeTagsResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListKnowledgeTagsResponseBodyData) SetTags(v []*ListKnowledgeTagsResponseBodyDataTags) *ListKnowledgeTagsResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListKnowledgeTagsResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeTagsResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// key_name
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListKnowledgeTagsResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeTagsResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListKnowledgeTagsResponseBodyDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListKnowledgeTagsResponseBodyDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListKnowledgeTagsResponseBodyDataTags) SetTagKey(v string) *ListKnowledgeTagsResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *ListKnowledgeTagsResponseBodyDataTags) SetTagValue(v string) *ListKnowledgeTagsResponseBodyDataTags {
	s.TagValue = &v
	return s
}

func (s *ListKnowledgeTagsResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
