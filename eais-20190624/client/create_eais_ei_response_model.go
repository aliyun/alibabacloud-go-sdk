// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaisEiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEaisEiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEaisEiResponse
	GetStatusCode() *int32
	SetBody(v *CreateEaisEiResponseBody) *CreateEaisEiResponse
	GetBody() *CreateEaisEiResponseBody
}

type CreateEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaisEiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEaisEiResponse) GoString() string {
	return s.String()
}

func (s *CreateEaisEiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEaisEiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEaisEiResponse) GetBody() *CreateEaisEiResponseBody {
	return s.Body
}

func (s *CreateEaisEiResponse) SetHeaders(v map[string]*string) *CreateEaisEiResponse {
	s.Headers = v
	return s
}

func (s *CreateEaisEiResponse) SetStatusCode(v int32) *CreateEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaisEiResponse) SetBody(v *CreateEaisEiResponseBody) *CreateEaisEiResponse {
	s.Body = v
	return s
}

func (s *CreateEaisEiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
