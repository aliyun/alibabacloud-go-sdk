// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportImageStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageExportStatus(v string) *DescribeExportImageStatusResponseBody
	GetImageExportStatus() *string
	SetRequestId(v string) *DescribeExportImageStatusResponseBody
	GetRequestId() *string
}

type DescribeExportImageStatusResponseBody struct {
	// The image export status. Valid values:
	//
	// - Exporting: The image is being exported.
	//
	// - Exported: The image is exported.
	//
	// - ExportError: An error occurred during the image export.
	//
	// - Unexported: The image is not exported.
	//
	// example:
	//
	// Exporting
	ImageExportStatus *string `json:"ImageExportStatus,omitempty" xml:"ImageExportStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14BBB3A0-3DBE-54F5-AEC8-01D3F6B1EBE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExportImageStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusResponseBody) GetImageExportStatus() *string {
	return s.ImageExportStatus
}

func (s *DescribeExportImageStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExportImageStatusResponseBody) SetImageExportStatus(v string) *DescribeExportImageStatusResponseBody {
	s.ImageExportStatus = &v
	return s
}

func (s *DescribeExportImageStatusResponseBody) SetRequestId(v string) *DescribeExportImageStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportImageStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
