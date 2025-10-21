// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlPreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSqlPreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSqlPreviewResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSqlPreviewResponseBody) *SubmitSqlPreviewResponse
	GetBody() *SubmitSqlPreviewResponseBody
}

type SubmitSqlPreviewResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSqlPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSqlPreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlPreviewResponse) GoString() string {
	return s.String()
}

func (s *SubmitSqlPreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSqlPreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSqlPreviewResponse) GetBody() *SubmitSqlPreviewResponseBody {
	return s.Body
}

func (s *SubmitSqlPreviewResponse) SetHeaders(v map[string]*string) *SubmitSqlPreviewResponse {
	s.Headers = v
	return s
}

func (s *SubmitSqlPreviewResponse) SetStatusCode(v int32) *SubmitSqlPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSqlPreviewResponse) SetBody(v *SubmitSqlPreviewResponseBody) *SubmitSqlPreviewResponse {
	s.Body = v
	return s
}

func (s *SubmitSqlPreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
