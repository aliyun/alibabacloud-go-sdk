// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentificationSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentificationSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentificationSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentificationSessionResponseBody) *GetIdentificationSessionResponse
	GetBody() *GetIdentificationSessionResponseBody
}

type GetIdentificationSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentificationSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentificationSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationSessionResponse) GoString() string {
	return s.String()
}

func (s *GetIdentificationSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentificationSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentificationSessionResponse) GetBody() *GetIdentificationSessionResponseBody {
	return s.Body
}

func (s *GetIdentificationSessionResponse) SetHeaders(v map[string]*string) *GetIdentificationSessionResponse {
	s.Headers = v
	return s
}

func (s *GetIdentificationSessionResponse) SetStatusCode(v int32) *GetIdentificationSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentificationSessionResponse) SetBody(v *GetIdentificationSessionResponseBody) *GetIdentificationSessionResponse {
	s.Body = v
	return s
}

func (s *GetIdentificationSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
