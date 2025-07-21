// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFbIssueLabelsByLCResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFbIssueLabelsByLCResponseBody
	GetCode() *string
	SetData(v *ListFbIssueLabelsByLCResponseBodyData) *ListFbIssueLabelsByLCResponseBody
	GetData() *ListFbIssueLabelsByLCResponseBodyData
	SetMessage(v string) *ListFbIssueLabelsByLCResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFbIssueLabelsByLCResponseBody
	GetRequestId() *string
}

type ListFbIssueLabelsByLCResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListFbIssueLabelsByLCResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFbIssueLabelsByLCResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsByLCResponseBody) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFbIssueLabelsByLCResponseBody) GetData() *ListFbIssueLabelsByLCResponseBodyData {
	return s.Data
}

func (s *ListFbIssueLabelsByLCResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFbIssueLabelsByLCResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFbIssueLabelsByLCResponseBody) SetCode(v string) *ListFbIssueLabelsByLCResponseBody {
	s.Code = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) SetData(v *ListFbIssueLabelsByLCResponseBodyData) *ListFbIssueLabelsByLCResponseBody {
	s.Data = v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) SetMessage(v string) *ListFbIssueLabelsByLCResponseBody {
	s.Message = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) SetRequestId(v string) *ListFbIssueLabelsByLCResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFbIssueLabelsByLCResponseBodyData struct {
	IssueLabel []*string `json:"IssueLabel,omitempty" xml:"IssueLabel,omitempty" type:"Repeated"`
}

func (s ListFbIssueLabelsByLCResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsByLCResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCResponseBodyData) GetIssueLabel() []*string {
	return s.IssueLabel
}

func (s *ListFbIssueLabelsByLCResponseBodyData) SetIssueLabel(v []*string) *ListFbIssueLabelsByLCResponseBodyData {
	s.IssueLabel = v
	return s
}

func (s *ListFbIssueLabelsByLCResponseBodyData) Validate() error {
	return dara.Validate(s)
}
