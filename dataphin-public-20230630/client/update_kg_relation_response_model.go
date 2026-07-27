// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKgRelationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKgRelationResponseBody) *UpdateKgRelationResponse
	GetBody() *UpdateKgRelationResponseBody
}

type UpdateKgRelationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgRelationResponse) GoString() string {
	return s.String()
}

func (s *UpdateKgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKgRelationResponse) GetBody() *UpdateKgRelationResponseBody {
	return s.Body
}

func (s *UpdateKgRelationResponse) SetHeaders(v map[string]*string) *UpdateKgRelationResponse {
	s.Headers = v
	return s
}

func (s *UpdateKgRelationResponse) SetStatusCode(v int32) *UpdateKgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKgRelationResponse) SetBody(v *UpdateKgRelationResponseBody) *UpdateKgRelationResponse {
	s.Body = v
	return s
}

func (s *UpdateKgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
