// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceTypePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceTypePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceTypePropertiesResponseBody) *GetResourceTypePropertiesResponse
	GetBody() *GetResourceTypePropertiesResponseBody
}

type GetResourceTypePropertiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceTypePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceTypePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypePropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceTypePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceTypePropertiesResponse) GetBody() *GetResourceTypePropertiesResponseBody {
	return s.Body
}

func (s *GetResourceTypePropertiesResponse) SetHeaders(v map[string]*string) *GetResourceTypePropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypePropertiesResponse) SetStatusCode(v int32) *GetResourceTypePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTypePropertiesResponse) SetBody(v *GetResourceTypePropertiesResponseBody) *GetResourceTypePropertiesResponse {
	s.Body = v
	return s
}

func (s *GetResourceTypePropertiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
