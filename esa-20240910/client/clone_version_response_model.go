// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneVersionResponse
	GetStatusCode() *int32
	SetBody(v *CloneVersionResponseBody) *CloneVersionResponse
	GetBody() *CloneVersionResponseBody
}

type CloneVersionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneVersionResponse) GoString() string {
	return s.String()
}

func (s *CloneVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneVersionResponse) GetBody() *CloneVersionResponseBody {
	return s.Body
}

func (s *CloneVersionResponse) SetHeaders(v map[string]*string) *CloneVersionResponse {
	s.Headers = v
	return s
}

func (s *CloneVersionResponse) SetStatusCode(v int32) *CloneVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneVersionResponse) SetBody(v *CloneVersionResponseBody) *CloneVersionResponse {
	s.Body = v
	return s
}

func (s *CloneVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
