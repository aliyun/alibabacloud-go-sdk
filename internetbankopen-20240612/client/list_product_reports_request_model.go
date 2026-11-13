// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLatestOnly(v bool) *ListProductReportsRequest
	GetLatestOnly() *bool
	SetProductCode(v string) *ListProductReportsRequest
	GetProductCode() *string
}

type ListProductReportsRequest struct {
	// Specifies whether to return only the most recent active epoch report. If this parameter is omitted or set to false, the operation returns active reports within the last ten completed natural weeks in UTC+8, up to ten records. If set to true, the operation queries the most recent active report from all completed historical epochs, returning at most one record without the ten-week limit. Only true or false is accepted. If a character string is passed, it must be strictly lowercase true or false. Other values return InvalidParameter.LatestOnly. This operation does not support paging.
	LatestOnly *bool `json:"LatestOnly,omitempty" xml:"LatestOnly,omitempty"`
	// The product code of the data product for which to query periodic reports. This must be consistent with the product code used when calling InvokeDataProduct. This parameter is required and allows only 1 to 128 ASCII letters, digits, periods (.), underscores (_), or hyphens (-). Leading and trailing spaces are not trimmed. If this parameter is missing or incorrectly formatted, InvalidParameter.ProductCode is returned. The query is limited to reports belonging to the current Alibaba Cloud account. If no available reports exist, an empty Reports array is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudsec.demo.v1
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductReportsRequest) GoString() string {
	return s.String()
}

func (s *ListProductReportsRequest) GetLatestOnly() *bool {
	return s.LatestOnly
}

func (s *ListProductReportsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductReportsRequest) SetLatestOnly(v bool) *ListProductReportsRequest {
	s.LatestOnly = &v
	return s
}

func (s *ListProductReportsRequest) SetProductCode(v string) *ListProductReportsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductReportsRequest) Validate() error {
	return dara.Validate(s)
}
