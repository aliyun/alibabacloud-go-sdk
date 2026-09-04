// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBasesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListVisibleKnowledgeBasesHeaders
	GetCommonHeaders() map[string]*string
	SetRequestId(v string) *ListVisibleKnowledgeBasesHeaders
	GetRequestId() *string
}

type ListVisibleKnowledgeBasesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F892C03F-7E12-5F37-A506-1FC3B065EAC6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListVisibleKnowledgeBasesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesHeaders) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListVisibleKnowledgeBasesHeaders) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVisibleKnowledgeBasesHeaders) SetCommonHeaders(v map[string]*string) *ListVisibleKnowledgeBasesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListVisibleKnowledgeBasesHeaders) SetRequestId(v string) *ListVisibleKnowledgeBasesHeaders {
	s.RequestId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesHeaders) Validate() error {
	return dara.Validate(s)
}
