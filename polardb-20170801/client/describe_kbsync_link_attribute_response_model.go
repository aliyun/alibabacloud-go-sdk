// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKBSyncLinkAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKBSyncLinkAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKBSyncLinkAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKBSyncLinkAttributeResponseBody) *DescribeKBSyncLinkAttributeResponse
	GetBody() *DescribeKBSyncLinkAttributeResponseBody
}

type DescribeKBSyncLinkAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKBSyncLinkAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKBSyncLinkAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinkAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinkAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKBSyncLinkAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKBSyncLinkAttributeResponse) GetBody() *DescribeKBSyncLinkAttributeResponseBody {
	return s.Body
}

func (s *DescribeKBSyncLinkAttributeResponse) SetHeaders(v map[string]*string) *DescribeKBSyncLinkAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponse) SetStatusCode(v int32) *DescribeKBSyncLinkAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponse) SetBody(v *DescribeKBSyncLinkAttributeResponseBody) *DescribeKBSyncLinkAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeKBSyncLinkAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
