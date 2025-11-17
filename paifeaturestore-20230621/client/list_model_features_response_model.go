// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelFeaturesResponseBody) *ListModelFeaturesResponse
	GetBody() *ListModelFeaturesResponseBody
}

type ListModelFeaturesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelFeaturesResponse) GetBody() *ListModelFeaturesResponseBody {
	return s.Body
}

func (s *ListModelFeaturesResponse) SetHeaders(v map[string]*string) *ListModelFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListModelFeaturesResponse) SetStatusCode(v int32) *ListModelFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelFeaturesResponse) SetBody(v *ListModelFeaturesResponseBody) *ListModelFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListModelFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
