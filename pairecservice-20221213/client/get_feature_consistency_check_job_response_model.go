// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConsistencyCheckJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFeatureConsistencyCheckJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFeatureConsistencyCheckJobResponse
	GetStatusCode() *int32
	SetBody(v *GetFeatureConsistencyCheckJobResponseBody) *GetFeatureConsistencyCheckJobResponse
	GetBody() *GetFeatureConsistencyCheckJobResponseBody
}

type GetFeatureConsistencyCheckJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureConsistencyCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureConsistencyCheckJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFeatureConsistencyCheckJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFeatureConsistencyCheckJobResponse) GetBody() *GetFeatureConsistencyCheckJobResponseBody {
	return s.Body
}

func (s *GetFeatureConsistencyCheckJobResponse) SetHeaders(v map[string]*string) *GetFeatureConsistencyCheckJobResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponse) SetStatusCode(v int32) *GetFeatureConsistencyCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponse) SetBody(v *GetFeatureConsistencyCheckJobResponseBody) *GetFeatureConsistencyCheckJobResponse {
	s.Body = v
	return s
}

func (s *GetFeatureConsistencyCheckJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
