// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSimQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSimQuestionResponseBody
	GetRequestId() *string
	SetSimQuestions(v []*ListSimQuestionResponseBodySimQuestions) *ListSimQuestionResponseBody
	GetSimQuestions() []*ListSimQuestionResponseBodySimQuestions
}

type ListSimQuestionResponseBody struct {
	// example:
	//
	// 15CD94CC-CBEB-4189-806C-A132D1F45D51
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimQuestions []*ListSimQuestionResponseBodySimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
}

func (s ListSimQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *ListSimQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSimQuestionResponseBody) GetSimQuestions() []*ListSimQuestionResponseBodySimQuestions {
	return s.SimQuestions
}

func (s *ListSimQuestionResponseBody) SetRequestId(v string) *ListSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSimQuestionResponseBody) SetSimQuestions(v []*ListSimQuestionResponseBodySimQuestions) *ListSimQuestionResponseBody {
	s.SimQuestions = v
	return s
}

func (s *ListSimQuestionResponseBody) Validate() error {
	if s.SimQuestions != nil {
		for _, item := range s.SimQuestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSimQuestionResponseBodySimQuestions struct {
	// example:
	//
	// 2022-05-30T02:08:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-13T03:49:28Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 30001979424
	SimQuestionId *int64  `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListSimQuestionResponseBodySimQuestions) String() string {
	return dara.Prettify(s)
}

func (s ListSimQuestionResponseBodySimQuestions) GoString() string {
	return s.String()
}

func (s *ListSimQuestionResponseBodySimQuestions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSimQuestionResponseBodySimQuestions) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListSimQuestionResponseBodySimQuestions) GetSimQuestionId() *int64 {
	return s.SimQuestionId
}

func (s *ListSimQuestionResponseBodySimQuestions) GetTitle() *string {
	return s.Title
}

func (s *ListSimQuestionResponseBodySimQuestions) SetCreateTime(v string) *ListSimQuestionResponseBodySimQuestions {
	s.CreateTime = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) SetModifyTime(v string) *ListSimQuestionResponseBodySimQuestions {
	s.ModifyTime = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) SetSimQuestionId(v int64) *ListSimQuestionResponseBodySimQuestions {
	s.SimQuestionId = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) SetTitle(v string) *ListSimQuestionResponseBodySimQuestions {
	s.Title = &v
	return s
}

func (s *ListSimQuestionResponseBodySimQuestions) Validate() error {
	return dara.Validate(s)
}
