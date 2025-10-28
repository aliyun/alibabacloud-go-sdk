// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListFeaturesResponseBody) *ListFeaturesResponse
	GetBody() *ListFeaturesResponseBody
}

type ListFeaturesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeaturesResponse) GetBody() *ListFeaturesResponseBody {
	return s.Body
}

func (s *ListFeaturesResponse) SetHeaders(v map[string]*string) *ListFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListFeaturesResponse) SetStatusCode(v int32) *ListFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeaturesResponse) SetBody(v *ListFeaturesResponseBody) *ListFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
