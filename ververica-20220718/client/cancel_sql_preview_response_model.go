// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSqlPreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSqlPreviewResponse
	GetStatusCode() *int32
	SetBody(v *CancelSqlPreviewResponseBody) *CancelSqlPreviewResponse
	GetBody() *CancelSqlPreviewResponseBody
}

type CancelSqlPreviewResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSqlPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSqlPreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPreviewResponse) GoString() string {
	return s.String()
}

func (s *CancelSqlPreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSqlPreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSqlPreviewResponse) GetBody() *CancelSqlPreviewResponseBody {
	return s.Body
}

func (s *CancelSqlPreviewResponse) SetHeaders(v map[string]*string) *CancelSqlPreviewResponse {
	s.Headers = v
	return s
}

func (s *CancelSqlPreviewResponse) SetStatusCode(v int32) *CancelSqlPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSqlPreviewResponse) SetBody(v *CancelSqlPreviewResponseBody) *CancelSqlPreviewResponse {
	s.Body = v
	return s
}

func (s *CancelSqlPreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
