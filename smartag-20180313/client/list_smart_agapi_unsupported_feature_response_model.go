// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartAGApiUnsupportedFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmartAGApiUnsupportedFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmartAGApiUnsupportedFeatureResponse
	GetStatusCode() *int32
	SetBody(v *ListSmartAGApiUnsupportedFeatureResponseBody) *ListSmartAGApiUnsupportedFeatureResponse
	GetBody() *ListSmartAGApiUnsupportedFeatureResponseBody
}

type ListSmartAGApiUnsupportedFeatureResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmartAGApiUnsupportedFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmartAGApiUnsupportedFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGApiUnsupportedFeatureResponse) GoString() string {
	return s.String()
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) GetBody() *ListSmartAGApiUnsupportedFeatureResponseBody {
	return s.Body
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) SetHeaders(v map[string]*string) *ListSmartAGApiUnsupportedFeatureResponse {
	s.Headers = v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) SetStatusCode(v int32) *ListSmartAGApiUnsupportedFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) SetBody(v *ListSmartAGApiUnsupportedFeatureResponseBody) *ListSmartAGApiUnsupportedFeatureResponse {
	s.Body = v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
