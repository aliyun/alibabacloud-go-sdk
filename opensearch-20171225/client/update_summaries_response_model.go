// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSummariesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSummariesResponseBody) *UpdateSummariesResponse
	GetBody() *UpdateSummariesResponseBody
}

type UpdateSummariesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSummariesResponse) GoString() string {
	return s.String()
}

func (s *UpdateSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSummariesResponse) GetBody() *UpdateSummariesResponseBody {
	return s.Body
}

func (s *UpdateSummariesResponse) SetHeaders(v map[string]*string) *UpdateSummariesResponse {
	s.Headers = v
	return s
}

func (s *UpdateSummariesResponse) SetStatusCode(v int32) *UpdateSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSummariesResponse) SetBody(v *UpdateSummariesResponseBody) *UpdateSummariesResponse {
	s.Body = v
	return s
}

func (s *UpdateSummariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
