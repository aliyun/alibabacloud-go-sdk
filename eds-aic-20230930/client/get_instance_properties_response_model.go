// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstancePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstancePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *GetInstancePropertiesResponseBody) *GetInstancePropertiesResponse
	GetBody() *GetInstancePropertiesResponseBody
}

type GetInstancePropertiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetInstancePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstancePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstancePropertiesResponse) GetBody() *GetInstancePropertiesResponseBody {
	return s.Body
}

func (s *GetInstancePropertiesResponse) SetHeaders(v map[string]*string) *GetInstancePropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetInstancePropertiesResponse) SetStatusCode(v int32) *GetInstancePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancePropertiesResponse) SetBody(v *GetInstancePropertiesResponseBody) *GetInstancePropertiesResponse {
	s.Body = v
	return s
}

func (s *GetInstancePropertiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
