// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAudienceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomAudienceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomAudienceResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomAudienceResponseBody) *CreateCustomAudienceResponse
	GetBody() *CreateCustomAudienceResponseBody
}

type CreateCustomAudienceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomAudienceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomAudienceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAudienceResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomAudienceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomAudienceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomAudienceResponse) GetBody() *CreateCustomAudienceResponseBody {
	return s.Body
}

func (s *CreateCustomAudienceResponse) SetHeaders(v map[string]*string) *CreateCustomAudienceResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomAudienceResponse) SetStatusCode(v int32) *CreateCustomAudienceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomAudienceResponse) SetBody(v *CreateCustomAudienceResponseBody) *CreateCustomAudienceResponse {
	s.Body = v
	return s
}

func (s *CreateCustomAudienceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
