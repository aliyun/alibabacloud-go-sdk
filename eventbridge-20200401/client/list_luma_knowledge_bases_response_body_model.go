// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaKnowledgeBasesResponseBody
	GetCode() *string
	SetData(v *ListLumaKnowledgeBasesResponseBodyData) *ListLumaKnowledgeBasesResponseBody
	GetData() *ListLumaKnowledgeBasesResponseBodyData
	SetMessage(v string) *ListLumaKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaKnowledgeBasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaKnowledgeBasesResponseBody
	GetSuccess() *bool
}

type ListLumaKnowledgeBasesResponseBody struct {
	// The response code. A value of Success indicates a successful call. Otherwise, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of knowledge bases bound to the agent. All results are returned at once without pagination.
	Data *ListLumaKnowledgeBasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket feedback.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaKnowledgeBasesResponseBody) GetData() *ListLumaKnowledgeBasesResponseBodyData {
	return s.Data
}

func (s *ListLumaKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaKnowledgeBasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaKnowledgeBasesResponseBody) SetCode(v string) *ListLumaKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetData(v *ListLumaKnowledgeBasesResponseBodyData) *ListLumaKnowledgeBasesResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetMessage(v string) *ListLumaKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetRequestId(v string) *ListLumaKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetSuccess(v bool) *ListLumaKnowledgeBasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaKnowledgeBasesResponseBodyData struct {
	// The list of knowledge bases bound to the agent.
	//
	// example:
	//
	// [{"KnowledgeBaseName":"my-knowledge-base"}]
	KnowledgeBases []*KnowledgeBase `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty" type:"Repeated"`
}

func (s ListLumaKnowledgeBasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesResponseBodyData) GetKnowledgeBases() []*KnowledgeBase {
	return s.KnowledgeBases
}

func (s *ListLumaKnowledgeBasesResponseBodyData) SetKnowledgeBases(v []*KnowledgeBase) *ListLumaKnowledgeBasesResponseBodyData {
	s.KnowledgeBases = v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBodyData) Validate() error {
	if s.KnowledgeBases != nil {
		for _, item := range s.KnowledgeBases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
