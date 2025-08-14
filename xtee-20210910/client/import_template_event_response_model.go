// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTemplateEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportTemplateEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportTemplateEventResponse
	GetStatusCode() *int32
	SetBody(v *ImportTemplateEventResponseBody) *ImportTemplateEventResponse
	GetBody() *ImportTemplateEventResponseBody
}

type ImportTemplateEventResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportTemplateEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportTemplateEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportTemplateEventResponse) GoString() string {
	return s.String()
}

func (s *ImportTemplateEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportTemplateEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportTemplateEventResponse) GetBody() *ImportTemplateEventResponseBody {
	return s.Body
}

func (s *ImportTemplateEventResponse) SetHeaders(v map[string]*string) *ImportTemplateEventResponse {
	s.Headers = v
	return s
}

func (s *ImportTemplateEventResponse) SetStatusCode(v int32) *ImportTemplateEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportTemplateEventResponse) SetBody(v *ImportTemplateEventResponseBody) *ImportTemplateEventResponse {
	s.Body = v
	return s
}

func (s *ImportTemplateEventResponse) Validate() error {
	return dara.Validate(s)
}
