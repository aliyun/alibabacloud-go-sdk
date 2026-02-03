// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDatasetDocumentResponseBody
	GetCode() *string
	SetData(v *GetDatasetDocumentResponseBodyData) *GetDatasetDocumentResponseBody
	GetData() *GetDatasetDocumentResponseBodyData
	SetHttpStatusCode(v int32) *GetDatasetDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDatasetDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDatasetDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatasetDocumentResponseBody
	GetSuccess() *bool
}

type GetDatasetDocumentResponseBody struct {
	// example:
	//
	// NoData
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDatasetDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatasetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDatasetDocumentResponseBody) GetData() *GetDatasetDocumentResponseBodyData {
	return s.Data
}

func (s *GetDatasetDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDatasetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDatasetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatasetDocumentResponseBody) SetCode(v string) *GetDatasetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetData(v *GetDatasetDocumentResponseBodyData) *GetDatasetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetHttpStatusCode(v int32) *GetDatasetDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetMessage(v string) *GetDatasetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetRequestId(v string) *GetDatasetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) SetSuccess(v bool) *GetDatasetDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatasetDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetDocumentResponseBodyData struct {
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	DisableHandleMultimodalMedia *bool `json:"DisableHandleMultimodalMedia,omitempty" xml:"DisableHandleMultimodalMedia,omitempty"`
	// example:
	//
	// 用户指定的文档唯一ID
	DocId   *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 内部文档唯一ID
	DocUuid  *string                                     `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	Extend1  *string                                     `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2  *string                                     `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3  *string                                     `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	Metadata *GetDatasetDocumentResponseBodyDataMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-14 08:54:33
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 来源
	SourceFrom *string `json:"SourceFrom,omitempty" xml:"SourceFrom,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tags    []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Title   *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDatasetDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBodyData) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *GetDatasetDocumentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetDatasetDocumentResponseBodyData) GetDisableHandleMultimodalMedia() *bool {
	return s.DisableHandleMultimodalMedia
}

func (s *GetDatasetDocumentResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *GetDatasetDocumentResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *GetDatasetDocumentResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *GetDatasetDocumentResponseBodyData) GetExtend1() *string {
	return s.Extend1
}

func (s *GetDatasetDocumentResponseBodyData) GetExtend2() *string {
	return s.Extend2
}

func (s *GetDatasetDocumentResponseBodyData) GetExtend3() *string {
	return s.Extend3
}

func (s *GetDatasetDocumentResponseBodyData) GetMetadata() *GetDatasetDocumentResponseBodyDataMetadata {
	return s.Metadata
}

func (s *GetDatasetDocumentResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *GetDatasetDocumentResponseBodyData) GetSourceFrom() *string {
	return s.SourceFrom
}

func (s *GetDatasetDocumentResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetDatasetDocumentResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetDatasetDocumentResponseBodyData) GetTags() []*string {
	return s.Tags
}

func (s *GetDatasetDocumentResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetDatasetDocumentResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDatasetDocumentResponseBodyData) SetCategoryUuid(v string) *GetDatasetDocumentResponseBodyData {
	s.CategoryUuid = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetContent(v string) *GetDatasetDocumentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDisableHandleMultimodalMedia(v bool) *GetDatasetDocumentResponseBodyData {
	s.DisableHandleMultimodalMedia = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDocId(v string) *GetDatasetDocumentResponseBodyData {
	s.DocId = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDocType(v string) *GetDatasetDocumentResponseBodyData {
	s.DocType = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetDocUuid(v string) *GetDatasetDocumentResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetExtend1(v string) *GetDatasetDocumentResponseBodyData {
	s.Extend1 = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetExtend2(v string) *GetDatasetDocumentResponseBodyData {
	s.Extend2 = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetExtend3(v string) *GetDatasetDocumentResponseBodyData {
	s.Extend3 = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetMetadata(v *GetDatasetDocumentResponseBodyDataMetadata) *GetDatasetDocumentResponseBodyData {
	s.Metadata = v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetPubTime(v string) *GetDatasetDocumentResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetSourceFrom(v string) *GetDatasetDocumentResponseBodyData {
	s.SourceFrom = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetStatus(v int32) *GetDatasetDocumentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetSummary(v string) *GetDatasetDocumentResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetTags(v []*string) *GetDatasetDocumentResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetTitle(v string) *GetDatasetDocumentResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) SetUrl(v string) *GetDatasetDocumentResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyData) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatasetDocumentResponseBodyDataMetadata struct {
	AsrSentences []*GetDatasetDocumentResponseBodyDataMetadataAsrSentences `json:"AsrSentences,omitempty" xml:"AsrSentences,omitempty" type:"Repeated"`
	KeyValues    []*GetDatasetDocumentResponseBodyDataMetadataKeyValues    `json:"KeyValues,omitempty" xml:"KeyValues,omitempty" type:"Repeated"`
	Text         *string                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
	VideoShots   []*GetDatasetDocumentResponseBodyDataMetadataVideoShots   `json:"VideoShots,omitempty" xml:"VideoShots,omitempty" type:"Repeated"`
}

func (s GetDatasetDocumentResponseBodyDataMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) GetAsrSentences() []*GetDatasetDocumentResponseBodyDataMetadataAsrSentences {
	return s.AsrSentences
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) GetKeyValues() []*GetDatasetDocumentResponseBodyDataMetadataKeyValues {
	return s.KeyValues
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) GetText() *string {
	return s.Text
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) GetVideoShots() []*GetDatasetDocumentResponseBodyDataMetadataVideoShots {
	return s.VideoShots
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) SetAsrSentences(v []*GetDatasetDocumentResponseBodyDataMetadataAsrSentences) *GetDatasetDocumentResponseBodyDataMetadata {
	s.AsrSentences = v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) SetKeyValues(v []*GetDatasetDocumentResponseBodyDataMetadataKeyValues) *GetDatasetDocumentResponseBodyDataMetadata {
	s.KeyValues = v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) SetText(v string) *GetDatasetDocumentResponseBodyDataMetadata {
	s.Text = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) SetVideoShots(v []*GetDatasetDocumentResponseBodyDataMetadataVideoShots) *GetDatasetDocumentResponseBodyDataMetadata {
	s.VideoShots = v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadata) Validate() error {
	if s.AsrSentences != nil {
		for _, item := range s.AsrSentences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.KeyValues != nil {
		for _, item := range s.KeyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoShots != nil {
		for _, item := range s.VideoShots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDatasetDocumentResponseBodyDataMetadataAsrSentences struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDatasetDocumentResponseBodyDataMetadataAsrSentences) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBodyDataMetadataAsrSentences) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) GetText() *string {
	return s.Text
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) SetEndTime(v int64) *GetDatasetDocumentResponseBodyDataMetadataAsrSentences {
	s.EndTime = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) SetStartTime(v int64) *GetDatasetDocumentResponseBodyDataMetadataAsrSentences {
	s.StartTime = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) SetText(v string) *GetDatasetDocumentResponseBodyDataMetadataAsrSentences {
	s.Text = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataAsrSentences) Validate() error {
	return dara.Validate(s)
}

type GetDatasetDocumentResponseBodyDataMetadataKeyValues struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDatasetDocumentResponseBodyDataMetadataKeyValues) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBodyDataMetadataKeyValues) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBodyDataMetadataKeyValues) GetKey() *string {
	return s.Key
}

func (s *GetDatasetDocumentResponseBodyDataMetadataKeyValues) GetValue() *string {
	return s.Value
}

func (s *GetDatasetDocumentResponseBodyDataMetadataKeyValues) SetKey(v string) *GetDatasetDocumentResponseBodyDataMetadataKeyValues {
	s.Key = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataKeyValues) SetValue(v string) *GetDatasetDocumentResponseBodyDataMetadataKeyValues {
	s.Value = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataKeyValues) Validate() error {
	return dara.Validate(s)
}

type GetDatasetDocumentResponseBodyDataMetadataVideoShots struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDatasetDocumentResponseBodyDataMetadataVideoShots) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentResponseBodyDataMetadataVideoShots) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) GetText() *string {
	return s.Text
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) SetEndTime(v int64) *GetDatasetDocumentResponseBodyDataMetadataVideoShots {
	s.EndTime = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) SetStartTime(v int64) *GetDatasetDocumentResponseBodyDataMetadataVideoShots {
	s.StartTime = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) SetText(v string) *GetDatasetDocumentResponseBodyDataMetadataVideoShots {
	s.Text = &v
	return s
}

func (s *GetDatasetDocumentResponseBodyDataMetadataVideoShots) Validate() error {
	return dara.Validate(s)
}
