// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnswerLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAnswerLibResponseBodyData) *ListAnswerLibResponseBody
	GetData() []*ListAnswerLibResponseBodyData
	SetRequestId(v string) *ListAnswerLibResponseBody
	GetRequestId() *string
}

type ListAnswerLibResponseBody struct {
	Data []*ListAnswerLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAnswerLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnswerLibResponseBody) GetData() []*ListAnswerLibResponseBodyData {
	return s.Data
}

func (s *ListAnswerLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnswerLibResponseBody) SetData(v []*ListAnswerLibResponseBodyData) *ListAnswerLibResponseBody {
	s.Data = v
	return s
}

func (s *ListAnswerLibResponseBody) SetRequestId(v string) *ListAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnswerLibResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnswerLibResponseBodyData struct {
	// example:
	//
	// 100
	AnswerCount *int32 `json:"AnswerCount,omitempty" xml:"AnswerCount,omitempty"`
	// example:
	//
	// 2024-06-03 18:15:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// alxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// UID。
	//
	// example:
	//
	// 1643953****74290
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListAnswerLibResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnswerLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnswerLibResponseBodyData) GetAnswerCount() *int32 {
	return s.AnswerCount
}

func (s *ListAnswerLibResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAnswerLibResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *ListAnswerLibResponseBodyData) GetLibName() *string {
	return s.LibName
}

func (s *ListAnswerLibResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *ListAnswerLibResponseBodyData) SetAnswerCount(v int32) *ListAnswerLibResponseBodyData {
	s.AnswerCount = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetGmtModified(v string) *ListAnswerLibResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetLibId(v string) *ListAnswerLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetLibName(v string) *ListAnswerLibResponseBodyData {
	s.LibName = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetUid(v string) *ListAnswerLibResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) Validate() error {
	return dara.Validate(s)
}
