// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantBookDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTextbookAssistantBookDirectoriesResponseBodyData) *ListTextbookAssistantBookDirectoriesResponseBody
	GetData() *ListTextbookAssistantBookDirectoriesResponseBodyData
	SetErrCode(v string) *ListTextbookAssistantBookDirectoriesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListTextbookAssistantBookDirectoriesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListTextbookAssistantBookDirectoriesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTextbookAssistantBookDirectoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTextbookAssistantBookDirectoriesResponseBody
	GetSuccess() *bool
}

type ListTextbookAssistantBookDirectoriesResponseBody struct {
	Data *ListTextbookAssistantBookDirectoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B_USER_NOT_FOUND_EXCEPTION
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A5E9849-A2F0-551D-A7D8-1A8118557BAB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) GetData() *ListTextbookAssistantBookDirectoriesResponseBodyData {
	return s.Data
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetData(v *ListTextbookAssistantBookDirectoriesResponseBodyData) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetErrCode(v string) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetErrMessage(v string) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetRequestId(v string) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetSuccess(v bool) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTextbookAssistantBookDirectoriesResponseBodyData struct {
	DirectoryTree []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree `json:"directoryTree,omitempty" xml:"directoryTree,omitempty" type:"Repeated"`
	EditionInfo   *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo     `json:"editionInfo,omitempty" xml:"editionInfo,omitempty" type:"Struct"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) GetDirectoryTree() []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	return s.DirectoryTree
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) GetEditionInfo() *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	return s.EditionInfo
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) SetDirectoryTree(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) *ListTextbookAssistantBookDirectoriesResponseBodyData {
	s.DirectoryTree = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) SetEditionInfo(v *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) *ListTextbookAssistantBookDirectoriesResponseBodyData {
	s.EditionInfo = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) Validate() error {
	if s.DirectoryTree != nil {
		for _, item := range s.DirectoryTree {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EditionInfo != nil {
		if err := s.EditionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree struct {
	// example:
	//
	// 05758807ed8e11eebe6e0c42a106bb02
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// example:
	//
	// 2 Jobs
	DirectoryName *string                                                                   `json:"directoryName,omitempty" xml:"directoryName,omitempty"`
	Topic         []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic `json:"topic,omitempty" xml:"topic,omitempty" type:"Repeated"`
	Unit          []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit  `json:"unit,omitempty" xml:"unit,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) GetTopic() []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic {
	return s.Topic
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) GetUnit() []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	return s.Unit
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetDirectoryId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetDirectoryName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.DirectoryName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetTopic(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetUnit(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.Unit = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) Validate() error {
	if s.Topic != nil {
		for _, item := range s.Topic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Unit != nil {
		for _, item := range s.Unit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic struct {
	// example:
	//
	// 1323
	LabelId   *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) GetLabelId() *string {
	return s.LabelId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) GetLabelName() *string {
	return s.LabelName
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) SetLabelId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic {
	s.LabelId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) SetLabelName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic {
	s.LabelName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit struct {
	DirectoryId   *string                                                                         `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	DirectoryName *string                                                                         `json:"directoryName,omitempty" xml:"directoryName,omitempty"`
	Section       []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection `json:"section,omitempty" xml:"section,omitempty" type:"Repeated"`
	Topic         []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic   `json:"topic,omitempty" xml:"topic,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) GetSection() []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	return s.Section
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) GetTopic() []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic {
	return s.Topic
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetDirectoryId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetDirectoryName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.DirectoryName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetSection(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.Section = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetTopic(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) Validate() error {
	if s.Section != nil {
		for _, item := range s.Section {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Topic != nil {
		for _, item := range s.Topic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection struct {
	Children      interface{}                                                                          `json:"children,omitempty" xml:"children,omitempty"`
	DirectoryId   *string                                                                              `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	DirectoryName *string                                                                              `json:"directoryName,omitempty" xml:"directoryName,omitempty"`
	Topic         []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic `json:"topic,omitempty" xml:"topic,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) GetChildren() interface{} {
	return s.Children
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) GetTopic() []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic {
	return s.Topic
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetChildren(v interface{}) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.Children = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetDirectoryId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetDirectoryName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.DirectoryName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetTopic(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) Validate() error {
	if s.Topic != nil {
		for _, item := range s.Topic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic struct {
	LabelId   *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) GetLabelId() *string {
	return s.LabelId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) GetLabelName() *string {
	return s.LabelName
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) SetLabelId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic {
	s.LabelId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) SetLabelName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic {
	s.LabelName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic struct {
	LabelId   *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) GetLabelId() *string {
	return s.LabelId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) GetLabelName() *string {
	return s.LabelName
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) SetLabelId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic {
	s.LabelId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) SetLabelName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic {
	s.LabelName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo struct {
	// example:
	//
	// 55857
	BookId *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	// example:
	//
	// 1
	BookVolume *string `json:"bookVolume,omitempty" xml:"bookVolume,omitempty"`
	// example:
	//
	// 2010-1(2)
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// example:
	//
	// 3
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 2019-1(10)
	Impression *string `json:"impression,omitempty" xml:"impression,omitempty"`
	// example:
	//
	// 9787544413695
	Isbn      *string `json:"isbn,omitempty" xml:"isbn,omitempty"`
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
	// example:
	//
	// ENGLISH
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetBookId() *string {
	return s.BookId
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetBookVolume() *string {
	return s.BookVolume
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetEdition() *string {
	return s.Edition
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetGrade() *string {
	return s.Grade
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetImpression() *string {
	return s.Impression
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetIsbn() *string {
	return s.Isbn
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetPublisher() *string {
	return s.Publisher
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetSubject() *string {
	return s.Subject
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GetVersion() *string {
	return s.Version
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetBookId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetBookVolume(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.BookVolume = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetEdition(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Edition = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetGrade(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetImpression(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Impression = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetIsbn(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Isbn = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetPublisher(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Publisher = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetSubject(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Subject = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetVersion(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Version = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) Validate() error {
	return dara.Validate(s)
}
