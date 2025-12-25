// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMosaicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMosaicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMosaicsResponse
	GetStatusCode() *int32
	SetBody(v *AddMosaicsResponseBody) *AddMosaicsResponse
	GetBody() *AddMosaicsResponseBody
}

type AddMosaicsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMosaicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMosaicsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMosaicsResponse) GoString() string {
	return s.String()
}

func (s *AddMosaicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMosaicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMosaicsResponse) GetBody() *AddMosaicsResponseBody {
	return s.Body
}

func (s *AddMosaicsResponse) SetHeaders(v map[string]*string) *AddMosaicsResponse {
	s.Headers = v
	return s
}

func (s *AddMosaicsResponse) SetStatusCode(v int32) *AddMosaicsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMosaicsResponse) SetBody(v *AddMosaicsResponseBody) *AddMosaicsResponse {
	s.Body = v
	return s
}

func (s *AddMosaicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
