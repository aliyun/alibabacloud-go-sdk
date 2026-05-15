// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityProjectResponseBody) *CreateQualityProjectResponse
	GetBody() *CreateQualityProjectResponseBody
}

type CreateQualityProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityProjectResponse) GetBody() *CreateQualityProjectResponseBody {
	return s.Body
}

func (s *CreateQualityProjectResponse) SetHeaders(v map[string]*string) *CreateQualityProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityProjectResponse) SetStatusCode(v int32) *CreateQualityProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityProjectResponse) SetBody(v *CreateQualityProjectResponseBody) *CreateQualityProjectResponse {
	s.Body = v
	return s
}

func (s *CreateQualityProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
