// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgRelationResponse
	GetStatusCode() *int32
	SetBody(v *GetAgRelationResponseBody) *GetAgRelationResponse
	GetBody() *GetAgRelationResponseBody
}

type GetAgRelationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgRelationResponse) GoString() string {
	return s.String()
}

func (s *GetAgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgRelationResponse) GetBody() *GetAgRelationResponseBody {
	return s.Body
}

func (s *GetAgRelationResponse) SetHeaders(v map[string]*string) *GetAgRelationResponse {
	s.Headers = v
	return s
}

func (s *GetAgRelationResponse) SetStatusCode(v int32) *GetAgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgRelationResponse) SetBody(v *GetAgRelationResponseBody) *GetAgRelationResponse {
	s.Body = v
	return s
}

func (s *GetAgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
