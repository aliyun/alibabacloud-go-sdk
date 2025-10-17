// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppRecordTemplateResponseBody) *DeleteAppRecordTemplateResponse
	GetBody() *DeleteAppRecordTemplateResponseBody
}

type DeleteAppRecordTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppRecordTemplateResponse) GetBody() *DeleteAppRecordTemplateResponseBody {
	return s.Body
}

func (s *DeleteAppRecordTemplateResponse) SetHeaders(v map[string]*string) *DeleteAppRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppRecordTemplateResponse) SetStatusCode(v int32) *DeleteAppRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppRecordTemplateResponse) SetBody(v *DeleteAppRecordTemplateResponseBody) *DeleteAppRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAppRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
