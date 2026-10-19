// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAccountAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAccountAttributesResponse
	GetStatusCode() *int32
	SetBody(v *SetAccountAttributesResponseBody) *SetAccountAttributesResponse
	GetBody() *SetAccountAttributesResponseBody
}

type SetAccountAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccountAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccountAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAccountAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetAccountAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAccountAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAccountAttributesResponse) GetBody() *SetAccountAttributesResponseBody {
	return s.Body
}

func (s *SetAccountAttributesResponse) SetHeaders(v map[string]*string) *SetAccountAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetAccountAttributesResponse) SetStatusCode(v int32) *SetAccountAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccountAttributesResponse) SetBody(v *SetAccountAttributesResponseBody) *SetAccountAttributesResponse {
	s.Body = v
	return s
}

func (s *SetAccountAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
