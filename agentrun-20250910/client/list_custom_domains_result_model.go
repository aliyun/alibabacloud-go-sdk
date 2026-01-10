// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomDomainsResult
	GetCode() *string
	SetData(v *ListCustomDomainsOutput) *ListCustomDomainsResult
	GetData() *ListCustomDomainsOutput
	SetRequestId(v string) *ListCustomDomainsResult
	GetRequestId() *string
}

type ListCustomDomainsResult struct {
	Code      *string                  `json:"code,omitempty" xml:"code,omitempty"`
	Data      *ListCustomDomainsOutput `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCustomDomainsResult) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsResult) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResult) GetCode() *string {
	return s.Code
}

func (s *ListCustomDomainsResult) GetData() *ListCustomDomainsOutput {
	return s.Data
}

func (s *ListCustomDomainsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomDomainsResult) SetCode(v string) *ListCustomDomainsResult {
	s.Code = &v
	return s
}

func (s *ListCustomDomainsResult) SetData(v *ListCustomDomainsOutput) *ListCustomDomainsResult {
	s.Data = v
	return s
}

func (s *ListCustomDomainsResult) SetRequestId(v string) *ListCustomDomainsResult {
	s.RequestId = &v
	return s
}

func (s *ListCustomDomainsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
