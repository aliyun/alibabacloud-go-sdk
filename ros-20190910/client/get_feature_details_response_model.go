// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFeatureDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFeatureDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetFeatureDetailsResponseBody) *GetFeatureDetailsResponse
	GetBody() *GetFeatureDetailsResponseBody
}

type GetFeatureDetailsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFeatureDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFeatureDetailsResponse) GetBody() *GetFeatureDetailsResponseBody {
	return s.Body
}

func (s *GetFeatureDetailsResponse) SetHeaders(v map[string]*string) *GetFeatureDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureDetailsResponse) SetStatusCode(v int32) *GetFeatureDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureDetailsResponse) SetBody(v *GetFeatureDetailsResponseBody) *GetFeatureDetailsResponse {
	s.Body = v
	return s
}

func (s *GetFeatureDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
