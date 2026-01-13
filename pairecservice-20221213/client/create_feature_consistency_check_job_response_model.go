// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureConsistencyCheckJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFeatureConsistencyCheckJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFeatureConsistencyCheckJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateFeatureConsistencyCheckJobResponseBody) *CreateFeatureConsistencyCheckJobResponse
	GetBody() *CreateFeatureConsistencyCheckJobResponseBody
}

type CreateFeatureConsistencyCheckJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeatureConsistencyCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFeatureConsistencyCheckJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFeatureConsistencyCheckJobResponse) GetBody() *CreateFeatureConsistencyCheckJobResponseBody {
	return s.Body
}

func (s *CreateFeatureConsistencyCheckJobResponse) SetHeaders(v map[string]*string) *CreateFeatureConsistencyCheckJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponse) SetStatusCode(v int32) *CreateFeatureConsistencyCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponse) SetBody(v *CreateFeatureConsistencyCheckJobResponseBody) *CreateFeatureConsistencyCheckJobResponse {
	s.Body = v
	return s
}

func (s *CreateFeatureConsistencyCheckJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
