// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemeTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSchemeTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSchemeTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSchemeTaskConfigResponseBody) *UpdateSchemeTaskConfigResponse
	GetBody() *UpdateSchemeTaskConfigResponseBody
}

type UpdateSchemeTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSchemeTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSchemeTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSchemeTaskConfigResponse) GetBody() *UpdateSchemeTaskConfigResponseBody {
	return s.Body
}

func (s *UpdateSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *UpdateSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSchemeTaskConfigResponse) SetStatusCode(v int32) *UpdateSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSchemeTaskConfigResponse) SetBody(v *UpdateSchemeTaskConfigResponseBody) *UpdateSchemeTaskConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateSchemeTaskConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
