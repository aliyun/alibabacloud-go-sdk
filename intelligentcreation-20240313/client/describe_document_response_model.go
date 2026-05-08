// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDocumentResponse
	GetStatusCode() *int32
	SetBody(v *DocumentResult) *DescribeDocumentResponse
	GetBody() *DocumentResult
}

type DescribeDocumentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocumentResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocumentResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDocumentResponse) GetBody() *DocumentResult {
	return s.Body
}

func (s *DescribeDocumentResponse) SetHeaders(v map[string]*string) *DescribeDocumentResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocumentResponse) SetStatusCode(v int32) *DescribeDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocumentResponse) SetBody(v *DocumentResult) *DescribeDocumentResponse {
	s.Body = v
	return s
}

func (s *DescribeDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
