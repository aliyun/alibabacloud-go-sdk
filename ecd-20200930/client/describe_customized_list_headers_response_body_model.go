// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedListHeadersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v []*DescribeCustomizedListHeadersResponseBodyHeaders) *DescribeCustomizedListHeadersResponseBody
	GetHeaders() []*DescribeCustomizedListHeadersResponseBodyHeaders
	SetRequestId(v string) *DescribeCustomizedListHeadersResponseBody
	GetRequestId() *string
}

type DescribeCustomizedListHeadersResponseBody struct {
	// The list of table headers.
	Headers []*DescribeCustomizedListHeadersResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomizedListHeadersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedListHeadersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedListHeadersResponseBody) GetHeaders() []*DescribeCustomizedListHeadersResponseBodyHeaders {
	return s.Headers
}

func (s *DescribeCustomizedListHeadersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizedListHeadersResponseBody) SetHeaders(v []*DescribeCustomizedListHeadersResponseBodyHeaders) *DescribeCustomizedListHeadersResponseBody {
	s.Headers = v
	return s
}

func (s *DescribeCustomizedListHeadersResponseBody) SetRequestId(v string) *DescribeCustomizedListHeadersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizedListHeadersResponseBody) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomizedListHeadersResponseBodyHeaders struct {
	// The display type of the table header. You can invoke [ModifyCustomizedListHeaders](https://help.aliyun.com/document_detail/452324.html) to change it.
	//
	// example:
	//
	// display
	DisplayType *string `json:"DisplayType,omitempty" xml:"DisplayType,omitempty"`
	// The key of the table header.
	//
	// example:
	//
	// system_data_disk
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	// The name of the table header.
	//
	// example:
	//
	// System disk/Data disk
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
}

func (s DescribeCustomizedListHeadersResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedListHeadersResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) SetDisplayType(v string) *DescribeCustomizedListHeadersResponseBodyHeaders {
	s.DisplayType = &v
	return s
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) SetHeaderKey(v string) *DescribeCustomizedListHeadersResponseBodyHeaders {
	s.HeaderKey = &v
	return s
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) SetHeaderName(v string) *DescribeCustomizedListHeadersResponseBodyHeaders {
	s.HeaderName = &v
	return s
}

func (s *DescribeCustomizedListHeadersResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}
