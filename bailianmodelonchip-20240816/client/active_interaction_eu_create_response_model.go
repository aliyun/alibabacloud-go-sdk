// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveInteractionEuCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveInteractionEuCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveInteractionEuCreateResponse
	GetStatusCode() *int32
	SetBody(v *ActiveInteractionEuCreateResponseBody) *ActiveInteractionEuCreateResponse
	GetBody() *ActiveInteractionEuCreateResponseBody
}

type ActiveInteractionEuCreateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveInteractionEuCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveInteractionEuCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionEuCreateResponse) GoString() string {
	return s.String()
}

func (s *ActiveInteractionEuCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveInteractionEuCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveInteractionEuCreateResponse) GetBody() *ActiveInteractionEuCreateResponseBody {
	return s.Body
}

func (s *ActiveInteractionEuCreateResponse) SetHeaders(v map[string]*string) *ActiveInteractionEuCreateResponse {
	s.Headers = v
	return s
}

func (s *ActiveInteractionEuCreateResponse) SetStatusCode(v int32) *ActiveInteractionEuCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveInteractionEuCreateResponse) SetBody(v *ActiveInteractionEuCreateResponseBody) *ActiveInteractionEuCreateResponse {
	s.Body = v
	return s
}

func (s *ActiveInteractionEuCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
