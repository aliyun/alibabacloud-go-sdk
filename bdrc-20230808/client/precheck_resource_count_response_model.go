// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckResourceCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrecheckResourceCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrecheckResourceCountResponse
	GetStatusCode() *int32
	SetBody(v *PrecheckResourceCountResponseBody) *PrecheckResourceCountResponse
	GetBody() *PrecheckResourceCountResponseBody
}

type PrecheckResourceCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckResourceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckResourceCountResponse) String() string {
	return dara.Prettify(s)
}

func (s PrecheckResourceCountResponse) GoString() string {
	return s.String()
}

func (s *PrecheckResourceCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrecheckResourceCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrecheckResourceCountResponse) GetBody() *PrecheckResourceCountResponseBody {
	return s.Body
}

func (s *PrecheckResourceCountResponse) SetHeaders(v map[string]*string) *PrecheckResourceCountResponse {
	s.Headers = v
	return s
}

func (s *PrecheckResourceCountResponse) SetStatusCode(v int32) *PrecheckResourceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckResourceCountResponse) SetBody(v *PrecheckResourceCountResponseBody) *PrecheckResourceCountResponse {
	s.Body = v
	return s
}

func (s *PrecheckResourceCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
