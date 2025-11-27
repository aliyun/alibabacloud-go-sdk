// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizTracesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*BizTraceConfig) *ListBizTracesResponseBody
	GetItems() []*BizTraceConfig
	SetMaxResults(v int32) *ListBizTracesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListBizTracesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBizTracesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBizTracesResponseBody
	GetTotalCount() *int32
}

type ListBizTracesResponseBody struct {
	Items []*BizTraceConfig `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aa9d0e569b88098a0e3155c29b473201a
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0CEC5375-C554-562B-A65F-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 66
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListBizTracesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBizTracesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBizTracesResponseBody) GetItems() []*BizTraceConfig {
	return s.Items
}

func (s *ListBizTracesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBizTracesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBizTracesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBizTracesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBizTracesResponseBody) SetItems(v []*BizTraceConfig) *ListBizTracesResponseBody {
	s.Items = v
	return s
}

func (s *ListBizTracesResponseBody) SetMaxResults(v int32) *ListBizTracesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBizTracesResponseBody) SetNextToken(v string) *ListBizTracesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBizTracesResponseBody) SetRequestId(v string) *ListBizTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBizTracesResponseBody) SetTotalCount(v int32) *ListBizTracesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBizTracesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
