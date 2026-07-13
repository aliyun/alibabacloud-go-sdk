// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *QueryFeaturesResponseBody) *QueryFeaturesResponse
	GetBody() *QueryFeaturesResponseBody
}

type QueryFeaturesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFeaturesResponse) GoString() string {
	return s.String()
}

func (s *QueryFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFeaturesResponse) GetBody() *QueryFeaturesResponseBody {
	return s.Body
}

func (s *QueryFeaturesResponse) SetHeaders(v map[string]*string) *QueryFeaturesResponse {
	s.Headers = v
	return s
}

func (s *QueryFeaturesResponse) SetStatusCode(v int32) *QueryFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFeaturesResponse) SetBody(v *QueryFeaturesResponseBody) *QueryFeaturesResponse {
	s.Body = v
	return s
}

func (s *QueryFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
