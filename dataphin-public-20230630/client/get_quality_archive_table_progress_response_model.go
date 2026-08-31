// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityArchiveTableProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityArchiveTableProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityArchiveTableProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityArchiveTableProgressResponseBody) *GetQualityArchiveTableProgressResponse
	GetBody() *GetQualityArchiveTableProgressResponseBody
}

type GetQualityArchiveTableProgressResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityArchiveTableProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityArchiveTableProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityArchiveTableProgressResponse) GoString() string {
	return s.String()
}

func (s *GetQualityArchiveTableProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityArchiveTableProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityArchiveTableProgressResponse) GetBody() *GetQualityArchiveTableProgressResponseBody {
	return s.Body
}

func (s *GetQualityArchiveTableProgressResponse) SetHeaders(v map[string]*string) *GetQualityArchiveTableProgressResponse {
	s.Headers = v
	return s
}

func (s *GetQualityArchiveTableProgressResponse) SetStatusCode(v int32) *GetQualityArchiveTableProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityArchiveTableProgressResponse) SetBody(v *GetQualityArchiveTableProgressResponseBody) *GetQualityArchiveTableProgressResponse {
	s.Body = v
	return s
}

func (s *GetQualityArchiveTableProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
