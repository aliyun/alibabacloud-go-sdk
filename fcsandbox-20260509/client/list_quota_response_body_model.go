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
	Code      *string  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string  `json:"message,omitempty" xml:"message,omitempty"`
	NextToken *string  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Quotas    []*Quota `json:"quotas,omitempty" xml:"quotas,omitempty" type:"Repeated"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
