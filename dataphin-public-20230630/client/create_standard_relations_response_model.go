// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardRelationsResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardRelationsResponseBody) *CreateStandardRelationsResponse
	GetBody() *CreateStandardRelationsResponseBody
}

type CreateStandardRelationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardRelationsResponse) GetBody() *CreateStandardRelationsResponseBody {
	return s.Body
}

func (s *CreateStandardRelationsResponse) SetHeaders(v map[string]*string) *CreateStandardRelationsResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardRelationsResponse) SetStatusCode(v int32) *CreateStandardRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardRelationsResponse) SetBody(v *CreateStandardRelationsResponseBody) *CreateStandardRelationsResponse {
	s.Body = v
	return s
}

func (s *CreateStandardRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
