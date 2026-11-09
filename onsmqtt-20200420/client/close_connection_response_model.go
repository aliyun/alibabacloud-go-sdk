// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CloseConnectionResponseBody) *CloseConnectionResponse
	GetBody() *CloseConnectionResponseBody
}

type CloseConnectionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseConnectionResponse) GoString() string {
	return s.String()
}

func (s *CloseConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseConnectionResponse) GetBody() *CloseConnectionResponseBody {
	return s.Body
}

func (s *CloseConnectionResponse) SetHeaders(v map[string]*string) *CloseConnectionResponse {
	s.Headers = v
	return s
}

func (s *CloseConnectionResponse) SetStatusCode(v int32) *CloseConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseConnectionResponse) SetBody(v *CloseConnectionResponseBody) *CloseConnectionResponse {
	s.Body = v
	return s
}

func (s *CloseConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
