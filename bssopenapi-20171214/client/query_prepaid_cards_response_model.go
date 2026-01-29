// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrepaidCardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPrepaidCardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPrepaidCardsResponse
	GetStatusCode() *int32
	SetBody(v *QueryPrepaidCardsResponseBody) *QueryPrepaidCardsResponse
	GetBody() *QueryPrepaidCardsResponseBody
}

type QueryPrepaidCardsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPrepaidCardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPrepaidCardsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPrepaidCardsResponse) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPrepaidCardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPrepaidCardsResponse) GetBody() *QueryPrepaidCardsResponseBody {
	return s.Body
}

func (s *QueryPrepaidCardsResponse) SetHeaders(v map[string]*string) *QueryPrepaidCardsResponse {
	s.Headers = v
	return s
}

func (s *QueryPrepaidCardsResponse) SetStatusCode(v int32) *QueryPrepaidCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPrepaidCardsResponse) SetBody(v *QueryPrepaidCardsResponseBody) *QueryPrepaidCardsResponse {
	s.Body = v
	return s
}

func (s *QueryPrepaidCardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
