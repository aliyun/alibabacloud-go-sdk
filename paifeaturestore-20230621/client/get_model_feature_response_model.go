// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelFeatureResponse
	GetStatusCode() *int32
	SetBody(v *GetModelFeatureResponseBody) *GetModelFeatureResponse
	GetBody() *GetModelFeatureResponseBody
}

type GetModelFeatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *GetModelFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelFeatureResponse) GetBody() *GetModelFeatureResponseBody {
	return s.Body
}

func (s *GetModelFeatureResponse) SetHeaders(v map[string]*string) *GetModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *GetModelFeatureResponse) SetStatusCode(v int32) *GetModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelFeatureResponse) SetBody(v *GetModelFeatureResponseBody) *GetModelFeatureResponse {
	s.Body = v
	return s
}

func (s *GetModelFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
