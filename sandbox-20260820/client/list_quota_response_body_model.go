// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQuotaResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListQuotaResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListQuotaResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListQuotaResponseBody
	GetNextToken() *string
	SetQuotas(v []*Quota) *ListQuotaResponseBody
	GetQuotas() []*Quota
	SetRequestId(v string) *ListQuotaResponseBody
	GetRequestId() *string
}

type ListQuotaResponseBody struct {
	// The response code. The value is 200 for a successful request.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The value of maxResults specified in the request. This field is returned only if a positive integer is specified for maxResults in the request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. The value is success for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page, which can be used for subsequent queries.
	//
	// example:
	//
	// qxGrXje86XMrYQ51aJMy
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of quota configurations.
	Quotas []*Quota `json:"quotas" xml:"quotas" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQuotaResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQuotaResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaResponseBody) GetQuotas() []*Quota {
	return s.Quotas
}

func (s *ListQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaResponseBody) SetCode(v string) *ListQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *ListQuotaResponseBody) SetMaxResults(v int32) *ListQuotaResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaResponseBody) SetMessage(v string) *ListQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *ListQuotaResponseBody) SetNextToken(v string) *ListQuotaResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotaResponseBody) SetQuotas(v []*Quota) *ListQuotaResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotaResponseBody) SetRequestId(v string) *ListQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaResponseBody) Validate() error {
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
