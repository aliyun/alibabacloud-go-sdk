// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticViewNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSemanticViewNamesResponseBodyData) *ListSemanticViewNamesResponseBody
	GetData() []*ListSemanticViewNamesResponseBodyData
	SetRequestId(v string) *ListSemanticViewNamesResponseBody
	GetRequestId() *string
}

type ListSemanticViewNamesResponseBody struct {
	// The returned result data.
	Data []*ListSemanticViewNamesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSemanticViewNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticViewNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSemanticViewNamesResponseBody) GetData() []*ListSemanticViewNamesResponseBodyData {
	return s.Data
}

func (s *ListSemanticViewNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSemanticViewNamesResponseBody) SetData(v []*ListSemanticViewNamesResponseBodyData) *ListSemanticViewNamesResponseBody {
	s.Data = v
	return s
}

func (s *ListSemanticViewNamesResponseBody) SetRequestId(v string) *ListSemanticViewNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSemanticViewNamesResponseBody) Validate() error {
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

type ListSemanticViewNamesResponseBodyData struct {
	// The annotation of the semantic view.
	//
	// example:
	//
	// 这是一个定义销售额相关指标的视图
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the semantic view.
	//
	// example:
	//
	// revenue_analysis
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
	// The schema in which the semantic view resides.
	//
	// example:
	//
	// sales_db
	ViewSchema *string `json:"ViewSchema,omitempty" xml:"ViewSchema,omitempty"`
}

func (s ListSemanticViewNamesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticViewNamesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSemanticViewNamesResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *ListSemanticViewNamesResponseBodyData) GetViewName() *string {
	return s.ViewName
}

func (s *ListSemanticViewNamesResponseBodyData) GetViewSchema() *string {
	return s.ViewSchema
}

func (s *ListSemanticViewNamesResponseBodyData) SetComment(v string) *ListSemanticViewNamesResponseBodyData {
	s.Comment = &v
	return s
}

func (s *ListSemanticViewNamesResponseBodyData) SetViewName(v string) *ListSemanticViewNamesResponseBodyData {
	s.ViewName = &v
	return s
}

func (s *ListSemanticViewNamesResponseBodyData) SetViewSchema(v string) *ListSemanticViewNamesResponseBodyData {
	s.ViewSchema = &v
	return s
}

func (s *ListSemanticViewNamesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
