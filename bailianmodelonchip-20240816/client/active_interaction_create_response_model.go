// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveInteractionCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveInteractionCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveInteractionCreateResponse
	GetStatusCode() *int32
	SetBody(v *ActiveInteractionCreateResponseBody) *ActiveInteractionCreateResponse
	GetBody() *ActiveInteractionCreateResponseBody
}

type ActiveInteractionCreateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveInteractionCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveInteractionCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionCreateResponse) GoString() string {
	return s.String()
}

func (s *ActiveInteractionCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveInteractionCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveInteractionCreateResponse) GetBody() *ActiveInteractionCreateResponseBody {
	return s.Body
}

func (s *ActiveInteractionCreateResponse) SetHeaders(v map[string]*string) *ActiveInteractionCreateResponse {
	s.Headers = v
	return s
}

func (s *ActiveInteractionCreateResponse) SetStatusCode(v int32) *ActiveInteractionCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveInteractionCreateResponse) SetBody(v *ActiveInteractionCreateResponseBody) *ActiveInteractionCreateResponse {
	s.Body = v
	return s
}

func (s *ActiveInteractionCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
