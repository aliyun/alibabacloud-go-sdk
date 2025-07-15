// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSummaryDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DescribeSummaryDataRequest
	GetSecurityToken() *string
}

type DescribeSummaryDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSummaryDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSummaryDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSummaryDataRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSummaryDataRequest) SetSecurityToken(v string) *DescribeSummaryDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSummaryDataRequest) Validate() error {
	return dara.Validate(s)
}
