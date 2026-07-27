// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKgRelationResponse
	GetStatusCode() *int32
	SetBody(v *GetKgRelationResponseBody) *GetKgRelationResponse
	GetBody() *GetKgRelationResponseBody
}

type GetKgRelationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKgRelationResponse) GoString() string {
	return s.String()
}

func (s *GetKgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKgRelationResponse) GetBody() *GetKgRelationResponseBody {
	return s.Body
}

func (s *GetKgRelationResponse) SetHeaders(v map[string]*string) *GetKgRelationResponse {
	s.Headers = v
	return s
}

func (s *GetKgRelationResponse) SetStatusCode(v int32) *GetKgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKgRelationResponse) SetBody(v *GetKgRelationResponseBody) *GetKgRelationResponse {
	s.Body = v
	return s
}

func (s *GetKgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
