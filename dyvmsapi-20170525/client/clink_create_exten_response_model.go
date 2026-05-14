// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateExtenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkCreateExtenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkCreateExtenResponse
	GetStatusCode() *int32
	SetBody(v *ClinkCreateExtenResponseBody) *ClinkCreateExtenResponse
	GetBody() *ClinkCreateExtenResponseBody
}

type ClinkCreateExtenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkCreateExtenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkCreateExtenResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateExtenResponse) GoString() string {
	return s.String()
}

func (s *ClinkCreateExtenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkCreateExtenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkCreateExtenResponse) GetBody() *ClinkCreateExtenResponseBody {
	return s.Body
}

func (s *ClinkCreateExtenResponse) SetHeaders(v map[string]*string) *ClinkCreateExtenResponse {
	s.Headers = v
	return s
}

func (s *ClinkCreateExtenResponse) SetStatusCode(v int32) *ClinkCreateExtenResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkCreateExtenResponse) SetBody(v *ClinkCreateExtenResponseBody) *ClinkCreateExtenResponse {
	s.Body = v
	return s
}

func (s *ClinkCreateExtenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
