// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveRightsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveRightsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveRightsResponse
	GetStatusCode() *int32
	SetBody(v *ActiveRightsResponseBody) *ActiveRightsResponse
	GetBody() *ActiveRightsResponseBody
}

type ActiveRightsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveRightsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveRightsResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveRightsResponse) GoString() string {
	return s.String()
}

func (s *ActiveRightsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveRightsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveRightsResponse) GetBody() *ActiveRightsResponseBody {
	return s.Body
}

func (s *ActiveRightsResponse) SetHeaders(v map[string]*string) *ActiveRightsResponse {
	s.Headers = v
	return s
}

func (s *ActiveRightsResponse) SetStatusCode(v int32) *ActiveRightsResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveRightsResponse) SetBody(v *ActiveRightsResponseBody) *ActiveRightsResponse {
	s.Body = v
	return s
}

func (s *ActiveRightsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
