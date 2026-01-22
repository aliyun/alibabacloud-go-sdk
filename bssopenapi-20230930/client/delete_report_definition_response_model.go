// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReportDefinitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReportDefinitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReportDefinitionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReportDefinitionResponseBody) *DeleteReportDefinitionResponse
	GetBody() *DeleteReportDefinitionResponseBody
}

type DeleteReportDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReportDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReportDefinitionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReportDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteReportDefinitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReportDefinitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReportDefinitionResponse) GetBody() *DeleteReportDefinitionResponseBody {
	return s.Body
}

func (s *DeleteReportDefinitionResponse) SetHeaders(v map[string]*string) *DeleteReportDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteReportDefinitionResponse) SetStatusCode(v int32) *DeleteReportDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReportDefinitionResponse) SetBody(v *DeleteReportDefinitionResponseBody) *DeleteReportDefinitionResponse {
	s.Body = v
	return s
}

func (s *DeleteReportDefinitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
