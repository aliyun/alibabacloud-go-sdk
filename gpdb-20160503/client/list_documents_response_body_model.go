// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListDocumentsResponseBody
	GetCount() *int32
	SetItems(v *ListDocumentsResponseBodyItems) *ListDocumentsResponseBody
	GetItems() *ListDocumentsResponseBodyItems
	SetMessage(v string) *ListDocumentsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDocumentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDocumentsResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListDocumentsResponseBody
	GetStatus() *string
}

type ListDocumentsResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The queried documents.
	Items *ListDocumentsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListDocumentsResponseBody) GetItems() *ListDocumentsResponseBodyItems {
	return s.Items
}

func (s *ListDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDocumentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListDocumentsResponseBody) SetCount(v int32) *ListDocumentsResponseBody {
	s.Count = &v
	return s
}

func (s *ListDocumentsResponseBody) SetItems(v *ListDocumentsResponseBodyItems) *ListDocumentsResponseBody {
	s.Items = v
	return s
}

func (s *ListDocumentsResponseBody) SetMessage(v string) *ListDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocumentsResponseBody) SetNextToken(v string) *ListDocumentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDocumentsResponseBody) SetRequestId(v string) *ListDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentsResponseBody) SetStatus(v string) *ListDocumentsResponseBody {
	s.Status = &v
	return s
}

func (s *ListDocumentsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDocumentsResponseBodyItems struct {
	DocumentList []*ListDocumentsResponseBodyItemsDocumentList `json:"DocumentList,omitempty" xml:"DocumentList,omitempty" type:"Repeated"`
}

func (s ListDocumentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBodyItems) GetDocumentList() []*ListDocumentsResponseBodyItemsDocumentList {
	return s.DocumentList
}

func (s *ListDocumentsResponseBodyItems) SetDocumentList(v []*ListDocumentsResponseBodyItemsDocumentList) *ListDocumentsResponseBodyItems {
	s.DocumentList = v
	return s
}

func (s *ListDocumentsResponseBodyItems) Validate() error {
	if s.DocumentList != nil {
		for _, item := range s.DocumentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDocumentsResponseBodyItemsDocumentList struct {
	// The name of the document.
	//
	// example:
	//
	// music.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The source of the document.
	//
	// example:
	//
	// http://oss.xxx/music.txt
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListDocumentsResponseBodyItemsDocumentList) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsResponseBodyItemsDocumentList) GoString() string {
	return s.String()
}

func (s *ListDocumentsResponseBodyItemsDocumentList) GetFileName() *string {
	return s.FileName
}

func (s *ListDocumentsResponseBodyItemsDocumentList) GetSource() *string {
	return s.Source
}

func (s *ListDocumentsResponseBodyItemsDocumentList) SetFileName(v string) *ListDocumentsResponseBodyItemsDocumentList {
	s.FileName = &v
	return s
}

func (s *ListDocumentsResponseBodyItemsDocumentList) SetSource(v string) *ListDocumentsResponseBodyItemsDocumentList {
	s.Source = &v
	return s
}

func (s *ListDocumentsResponseBodyItemsDocumentList) Validate() error {
	return dara.Validate(s)
}
