// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIndividuationTextTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *QueryIndividuationTextTaskResponseBody
	GetCreateTime() *string
	SetRequestId(v string) *QueryIndividuationTextTaskResponseBody
	GetRequestId() *string
	SetStatus(v int32) *QueryIndividuationTextTaskResponseBody
	GetStatus() *int32
	SetTextList(v []*QueryIndividuationTextTaskResponseBodyTextList) *QueryIndividuationTextTaskResponseBody
	GetTextList() []*QueryIndividuationTextTaskResponseBodyTextList
	SetUpdateTime(v string) *QueryIndividuationTextTaskResponseBody
	GetUpdateTime() *string
}

type QueryIndividuationTextTaskResponseBody struct {
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 56AC346B-AF40-5E4F-AFFE-FD8BA5E6FB3A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	Status   *int32                                            `json:"status,omitempty" xml:"status,omitempty"`
	TextList []*QueryIndividuationTextTaskResponseBodyTextList `json:"textList,omitempty" xml:"textList,omitempty" type:"Repeated"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s QueryIndividuationTextTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryIndividuationTextTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryIndividuationTextTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryIndividuationTextTaskResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *QueryIndividuationTextTaskResponseBody) GetTextList() []*QueryIndividuationTextTaskResponseBodyTextList {
	return s.TextList
}

func (s *QueryIndividuationTextTaskResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryIndividuationTextTaskResponseBody) SetCreateTime(v string) *QueryIndividuationTextTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetRequestId(v string) *QueryIndividuationTextTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetStatus(v int32) *QueryIndividuationTextTaskResponseBody {
	s.Status = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetTextList(v []*QueryIndividuationTextTaskResponseBodyTextList) *QueryIndividuationTextTaskResponseBody {
	s.TextList = v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) SetUpdateTime(v string) *QueryIndividuationTextTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBody) Validate() error {
	if s.TextList != nil {
		for _, item := range s.TextList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryIndividuationTextTaskResponseBodyTextList struct {
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2761
	TextId *string `json:"textId,omitempty" xml:"textId,omitempty"`
	// example:
	//
	// 11
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryIndividuationTextTaskResponseBodyTextList) String() string {
	return dara.Prettify(s)
}

func (s QueryIndividuationTextTaskResponseBodyTextList) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) GetTextId() *string {
	return s.TextId
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) GetUserId() *string {
	return s.UserId
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) SetStatus(v int32) *QueryIndividuationTextTaskResponseBodyTextList {
	s.Status = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) SetTextId(v string) *QueryIndividuationTextTaskResponseBodyTextList {
	s.TextId = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) SetUserId(v string) *QueryIndividuationTextTaskResponseBodyTextList {
	s.UserId = &v
	return s
}

func (s *QueryIndividuationTextTaskResponseBodyTextList) Validate() error {
	return dara.Validate(s)
}
