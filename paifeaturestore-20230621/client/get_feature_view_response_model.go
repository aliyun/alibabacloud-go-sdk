// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFeatureViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFeatureViewResponse
	GetStatusCode() *int32
	SetBody(v *GetFeatureViewResponseBody) *GetFeatureViewResponse
	GetBody() *GetFeatureViewResponseBody
}

type GetFeatureViewResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureViewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureViewResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFeatureViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFeatureViewResponse) GetBody() *GetFeatureViewResponseBody {
	return s.Body
}

func (s *GetFeatureViewResponse) SetHeaders(v map[string]*string) *GetFeatureViewResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureViewResponse) SetStatusCode(v int32) *GetFeatureViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureViewResponse) SetBody(v *GetFeatureViewResponseBody) *GetFeatureViewResponse {
	s.Body = v
	return s
}

func (s *GetFeatureViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
