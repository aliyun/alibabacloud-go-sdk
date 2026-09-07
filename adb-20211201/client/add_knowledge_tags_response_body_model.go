// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddKnowledgeTagsResponseBodyData) *AddKnowledgeTagsResponseBody
	GetData() *AddKnowledgeTagsResponseBodyData
	SetRequestId(v string) *AddKnowledgeTagsResponseBody
	GetRequestId() *string
}

type AddKnowledgeTagsResponseBody struct {
	// The returned data.
	Data *AddKnowledgeTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddKnowledgeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddKnowledgeTagsResponseBody) GetData() *AddKnowledgeTagsResponseBodyData {
	return s.Data
}

func (s *AddKnowledgeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKnowledgeTagsResponseBody) SetData(v *AddKnowledgeTagsResponseBodyData) *AddKnowledgeTagsResponseBody {
	s.Data = v
	return s
}

func (s *AddKnowledgeTagsResponseBody) SetRequestId(v string) *AddKnowledgeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKnowledgeTagsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddKnowledgeTagsResponseBodyData struct {
	// The location of the knowledge base file.
	//
	// example:
	//
	// oss://bucketName/path/to/file.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The message.
	//
	// example:
	//
	// 1 tag skipped
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of tags that were successfully updated.
	//
	// example:
	//
	// 1
	Replaced *int32 `json:"Replaced,omitempty" xml:"Replaced,omitempty"`
	// The list of skipped tags.
	Skipped []*AddKnowledgeTagsResponseBodyDataSkipped `json:"Skipped,omitempty" xml:"Skipped,omitempty" type:"Repeated"`
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
	// The number of tags that were successfully added.
	//
	// example:
	//
	// 2
	Written *int32 `json:"Written,omitempty" xml:"Written,omitempty"`
}

func (s AddKnowledgeTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKnowledgeTagsResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *AddKnowledgeTagsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AddKnowledgeTagsResponseBodyData) GetReplaced() *int32 {
	return s.Replaced
}

func (s *AddKnowledgeTagsResponseBodyData) GetSkipped() []*AddKnowledgeTagsResponseBodyDataSkipped {
	return s.Skipped
}

func (s *AddKnowledgeTagsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddKnowledgeTagsResponseBodyData) GetWritten() *int32 {
	return s.Written
}

func (s *AddKnowledgeTagsResponseBodyData) SetFileLocation(v string) *AddKnowledgeTagsResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyData) SetMessage(v string) *AddKnowledgeTagsResponseBodyData {
	s.Message = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyData) SetReplaced(v int32) *AddKnowledgeTagsResponseBodyData {
	s.Replaced = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyData) SetSkipped(v []*AddKnowledgeTagsResponseBodyDataSkipped) *AddKnowledgeTagsResponseBodyData {
	s.Skipped = v
	return s
}

func (s *AddKnowledgeTagsResponseBodyData) SetSuccess(v bool) *AddKnowledgeTagsResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyData) SetWritten(v int32) *AddKnowledgeTagsResponseBodyData {
	s.Written = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyData) Validate() error {
	if s.Skipped != nil {
		for _, item := range s.Skipped {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddKnowledgeTagsResponseBodyDataSkipped struct {
	// The reason why the tag was skipped.
	//
	// example:
	//
	// conflict
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// skipKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// skipValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s AddKnowledgeTagsResponseBodyDataSkipped) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeTagsResponseBodyDataSkipped) GoString() string {
	return s.String()
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) GetReason() *string {
	return s.Reason
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) GetTagKey() *string {
	return s.TagKey
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) GetTagValue() *string {
	return s.TagValue
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) SetReason(v string) *AddKnowledgeTagsResponseBodyDataSkipped {
	s.Reason = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) SetTagKey(v string) *AddKnowledgeTagsResponseBodyDataSkipped {
	s.TagKey = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) SetTagValue(v string) *AddKnowledgeTagsResponseBodyDataSkipped {
	s.TagValue = &v
	return s
}

func (s *AddKnowledgeTagsResponseBodyDataSkipped) Validate() error {
	return dara.Validate(s)
}
