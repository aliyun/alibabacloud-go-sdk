// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTemplateDownloadResponseBody
	GetRequestId() *string
	SetData(v bool) *DescribeTemplateDownloadResponseBody
	GetData() *bool
}

type DescribeTemplateDownloadResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data object
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DescribeTemplateDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplateDownloadResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeTemplateDownloadResponseBody) SetRequestId(v string) *DescribeTemplateDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateDownloadResponseBody) SetData(v bool) *DescribeTemplateDownloadResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeTemplateDownloadResponseBody) Validate() error {
	return dara.Validate(s)
}
