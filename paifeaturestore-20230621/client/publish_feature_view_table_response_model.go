// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFeatureViewTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishFeatureViewTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishFeatureViewTableResponse
	GetStatusCode() *int32
	SetBody(v *PublishFeatureViewTableResponseBody) *PublishFeatureViewTableResponse
	GetBody() *PublishFeatureViewTableResponseBody
}

type PublishFeatureViewTableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishFeatureViewTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFeatureViewTableResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishFeatureViewTableResponse) GoString() string {
	return s.String()
}

func (s *PublishFeatureViewTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishFeatureViewTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishFeatureViewTableResponse) GetBody() *PublishFeatureViewTableResponseBody {
	return s.Body
}

func (s *PublishFeatureViewTableResponse) SetHeaders(v map[string]*string) *PublishFeatureViewTableResponse {
	s.Headers = v
	return s
}

func (s *PublishFeatureViewTableResponse) SetStatusCode(v int32) *PublishFeatureViewTableResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFeatureViewTableResponse) SetBody(v *PublishFeatureViewTableResponseBody) *PublishFeatureViewTableResponse {
	s.Body = v
	return s
}

func (s *PublishFeatureViewTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
