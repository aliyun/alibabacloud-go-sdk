// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHBaseSlbServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHBaseSlbServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHBaseSlbServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateHBaseSlbServerResponseBody) *CreateHBaseSlbServerResponse
	GetBody() *CreateHBaseSlbServerResponseBody
}

type CreateHBaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHBaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHBaseSlbServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHBaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHBaseSlbServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHBaseSlbServerResponse) GetBody() *CreateHBaseSlbServerResponseBody {
	return s.Body
}

func (s *CreateHBaseSlbServerResponse) SetHeaders(v map[string]*string) *CreateHBaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *CreateHBaseSlbServerResponse) SetStatusCode(v int32) *CreateHBaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHBaseSlbServerResponse) SetBody(v *CreateHBaseSlbServerResponseBody) *CreateHBaseSlbServerResponse {
	s.Body = v
	return s
}

func (s *CreateHBaseSlbServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
