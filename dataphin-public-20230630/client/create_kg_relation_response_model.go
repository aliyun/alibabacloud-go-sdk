// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKgRelationResponse
	GetStatusCode() *int32
	SetBody(v *CreateKgRelationResponseBody) *CreateKgRelationResponse
	GetBody() *CreateKgRelationResponseBody
}

type CreateKgRelationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateKgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKgRelationResponse) GetBody() *CreateKgRelationResponseBody {
	return s.Body
}

func (s *CreateKgRelationResponse) SetHeaders(v map[string]*string) *CreateKgRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateKgRelationResponse) SetStatusCode(v int32) *CreateKgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKgRelationResponse) SetBody(v *CreateKgRelationResponseBody) *CreateKgRelationResponse {
	s.Body = v
	return s
}

func (s *CreateKgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
