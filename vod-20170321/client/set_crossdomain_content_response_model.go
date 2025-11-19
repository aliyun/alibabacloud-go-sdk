// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCrossdomainContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCrossdomainContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCrossdomainContentResponse
	GetStatusCode() *int32
	SetBody(v *SetCrossdomainContentResponseBody) *SetCrossdomainContentResponse
	GetBody() *SetCrossdomainContentResponseBody
}

type SetCrossdomainContentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCrossdomainContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCrossdomainContentResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCrossdomainContentResponse) GoString() string {
	return s.String()
}

func (s *SetCrossdomainContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCrossdomainContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCrossdomainContentResponse) GetBody() *SetCrossdomainContentResponseBody {
	return s.Body
}

func (s *SetCrossdomainContentResponse) SetHeaders(v map[string]*string) *SetCrossdomainContentResponse {
	s.Headers = v
	return s
}

func (s *SetCrossdomainContentResponse) SetStatusCode(v int32) *SetCrossdomainContentResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCrossdomainContentResponse) SetBody(v *SetCrossdomainContentResponseBody) *SetCrossdomainContentResponse {
	s.Body = v
	return s
}

func (s *SetCrossdomainContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
