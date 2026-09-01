// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKBSyncLinksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKBSyncLinksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKBSyncLinksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKBSyncLinksResponseBody) *DescribeKBSyncLinksResponse
	GetBody() *DescribeKBSyncLinksResponseBody
}

type DescribeKBSyncLinksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKBSyncLinksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKBSyncLinksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinksResponse) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKBSyncLinksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKBSyncLinksResponse) GetBody() *DescribeKBSyncLinksResponseBody {
	return s.Body
}

func (s *DescribeKBSyncLinksResponse) SetHeaders(v map[string]*string) *DescribeKBSyncLinksResponse {
	s.Headers = v
	return s
}

func (s *DescribeKBSyncLinksResponse) SetStatusCode(v int32) *DescribeKBSyncLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKBSyncLinksResponse) SetBody(v *DescribeKBSyncLinksResponseBody) *DescribeKBSyncLinksResponse {
	s.Body = v
	return s
}

func (s *DescribeKBSyncLinksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
