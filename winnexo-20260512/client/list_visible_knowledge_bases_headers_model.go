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
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
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
