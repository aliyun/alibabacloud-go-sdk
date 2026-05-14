// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkUpdateClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkUpdateClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkUpdateClientResponse
	GetStatusCode() *int32
	SetBody(v *ClinkUpdateClientResponseBody) *ClinkUpdateClientResponse
	GetBody() *ClinkUpdateClientResponseBody
}

type ClinkUpdateClientResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkUpdateClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkUpdateClientResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientResponse) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkUpdateClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkUpdateClientResponse) GetBody() *ClinkUpdateClientResponseBody {
	return s.Body
}

func (s *ClinkUpdateClientResponse) SetHeaders(v map[string]*string) *ClinkUpdateClientResponse {
	s.Headers = v
	return s
}

func (s *ClinkUpdateClientResponse) SetStatusCode(v int32) *ClinkUpdateClientResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkUpdateClientResponse) SetBody(v *ClinkUpdateClientResponseBody) *ClinkUpdateClientResponse {
	s.Body = v
	return s
}

func (s *ClinkUpdateClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
