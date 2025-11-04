// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSearchIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropSearchIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropSearchIndexResponse
	GetStatusCode() *int32
	SetBody(v *DropSearchIndexResponseBody) *DropSearchIndexResponse
	GetBody() *DropSearchIndexResponseBody
}

type DropSearchIndexResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DropSearchIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DropSearchIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s DropSearchIndexResponse) GoString() string {
	return s.String()
}

func (s *DropSearchIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropSearchIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropSearchIndexResponse) GetBody() *DropSearchIndexResponseBody {
	return s.Body
}

func (s *DropSearchIndexResponse) SetHeaders(v map[string]*string) *DropSearchIndexResponse {
	s.Headers = v
	return s
}

func (s *DropSearchIndexResponse) SetStatusCode(v int32) *DropSearchIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DropSearchIndexResponse) SetBody(v *DropSearchIndexResponseBody) *DropSearchIndexResponse {
	s.Body = v
	return s
}

func (s *DropSearchIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
