// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalogCommonResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CatalogCommonResult
	GetRequestId() *string
	SetResult(v bool) *CatalogCommonResult
	GetResult() *bool
}

type CatalogCommonResult struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CatalogCommonResult) String() string {
	return dara.Prettify(s)
}

func (s CatalogCommonResult) GoString() string {
	return s.String()
}

func (s *CatalogCommonResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CatalogCommonResult) GetResult() *bool {
	return s.Result
}

func (s *CatalogCommonResult) SetRequestId(v string) *CatalogCommonResult {
	s.RequestId = &v
	return s
}

func (s *CatalogCommonResult) SetResult(v bool) *CatalogCommonResult {
	s.Result = &v
	return s
}

func (s *CatalogCommonResult) Validate() error {
	return dara.Validate(s)
}
