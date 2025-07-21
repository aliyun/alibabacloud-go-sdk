// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFbIssueLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFbIssueLabelsResponseBody
	GetCode() *string
	SetData(v *ListFbIssueLabelsResponseBodyData) *ListFbIssueLabelsResponseBody
	GetData() *ListFbIssueLabelsResponseBodyData
	SetMessage(v string) *ListFbIssueLabelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFbIssueLabelsResponseBody
	GetRequestId() *string
}

type ListFbIssueLabelsResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListFbIssueLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFbIssueLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFbIssueLabelsResponseBody) GetData() *ListFbIssueLabelsResponseBodyData {
	return s.Data
}

func (s *ListFbIssueLabelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFbIssueLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFbIssueLabelsResponseBody) SetCode(v string) *ListFbIssueLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFbIssueLabelsResponseBody) SetData(v *ListFbIssueLabelsResponseBodyData) *ListFbIssueLabelsResponseBody {
	s.Data = v
	return s
}

func (s *ListFbIssueLabelsResponseBody) SetMessage(v string) *ListFbIssueLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFbIssueLabelsResponseBody) SetRequestId(v string) *ListFbIssueLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFbIssueLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFbIssueLabelsResponseBodyData struct {
	IssueLabel []*string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty" type:"Repeated"`
}

func (s ListFbIssueLabelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsResponseBodyData) GetIssueLabel() []*string {
	return s.IssueLabel
}

func (s *ListFbIssueLabelsResponseBodyData) SetIssueLabel(v []*string) *ListFbIssueLabelsResponseBodyData {
	s.IssueLabel = v
	return s
}

func (s *ListFbIssueLabelsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
