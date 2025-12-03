// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHbaseSlbServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHbaseSlbServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHbaseSlbServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateHbaseSlbServerResponseBody) *CreateHbaseSlbServerResponse
	GetBody() *CreateHbaseSlbServerResponseBody
}

type CreateHbaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHbaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHbaseSlbServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHbaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *CreateHbaseSlbServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHbaseSlbServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHbaseSlbServerResponse) GetBody() *CreateHbaseSlbServerResponseBody {
	return s.Body
}

func (s *CreateHbaseSlbServerResponse) SetHeaders(v map[string]*string) *CreateHbaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *CreateHbaseSlbServerResponse) SetStatusCode(v int32) *CreateHbaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHbaseSlbServerResponse) SetBody(v *CreateHbaseSlbServerResponseBody) *CreateHbaseSlbServerResponse {
	s.Body = v
	return s
}

func (s *CreateHbaseSlbServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
