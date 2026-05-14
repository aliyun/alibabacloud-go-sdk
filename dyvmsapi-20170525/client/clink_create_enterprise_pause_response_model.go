// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateEnterprisePauseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkCreateEnterprisePauseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkCreateEnterprisePauseResponse
	GetStatusCode() *int32
	SetBody(v *ClinkCreateEnterprisePauseResponseBody) *ClinkCreateEnterprisePauseResponse
	GetBody() *ClinkCreateEnterprisePauseResponseBody
}

type ClinkCreateEnterprisePauseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkCreateEnterprisePauseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkCreateEnterprisePauseResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateEnterprisePauseResponse) GoString() string {
	return s.String()
}

func (s *ClinkCreateEnterprisePauseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkCreateEnterprisePauseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkCreateEnterprisePauseResponse) GetBody() *ClinkCreateEnterprisePauseResponseBody {
	return s.Body
}

func (s *ClinkCreateEnterprisePauseResponse) SetHeaders(v map[string]*string) *ClinkCreateEnterprisePauseResponse {
	s.Headers = v
	return s
}

func (s *ClinkCreateEnterprisePauseResponse) SetStatusCode(v int32) *ClinkCreateEnterprisePauseResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponse) SetBody(v *ClinkCreateEnterprisePauseResponseBody) *ClinkCreateEnterprisePauseResponse {
	s.Body = v
	return s
}

func (s *ClinkCreateEnterprisePauseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
