// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrailResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrailResponseBody) *CreateTrailResponse
	GetBody() *CreateTrailResponseBody
}

type CreateTrailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrailResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrailResponse) GoString() string {
	return s.String()
}

func (s *CreateTrailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrailResponse) GetBody() *CreateTrailResponseBody {
	return s.Body
}

func (s *CreateTrailResponse) SetHeaders(v map[string]*string) *CreateTrailResponse {
	s.Headers = v
	return s
}

func (s *CreateTrailResponse) SetStatusCode(v int32) *CreateTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrailResponse) SetBody(v *CreateTrailResponseBody) *CreateTrailResponse {
	s.Body = v
	return s
}

func (s *CreateTrailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
