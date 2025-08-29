// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackflowFeatureConsistencyCheckJobDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BackflowFeatureConsistencyCheckJobDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BackflowFeatureConsistencyCheckJobDataResponse
	GetStatusCode() *int32
	SetBody(v *BackflowFeatureConsistencyCheckJobDataResponseBody) *BackflowFeatureConsistencyCheckJobDataResponse
	GetBody() *BackflowFeatureConsistencyCheckJobDataResponseBody
}

type BackflowFeatureConsistencyCheckJobDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackflowFeatureConsistencyCheckJobDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackflowFeatureConsistencyCheckJobDataResponse) String() string {
	return dara.Prettify(s)
}

func (s BackflowFeatureConsistencyCheckJobDataResponse) GoString() string {
	return s.String()
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) GetBody() *BackflowFeatureConsistencyCheckJobDataResponseBody {
	return s.Body
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) SetHeaders(v map[string]*string) *BackflowFeatureConsistencyCheckJobDataResponse {
	s.Headers = v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) SetStatusCode(v int32) *BackflowFeatureConsistencyCheckJobDataResponse {
	s.StatusCode = &v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) SetBody(v *BackflowFeatureConsistencyCheckJobDataResponseBody) *BackflowFeatureConsistencyCheckJobDataResponse {
	s.Body = v
	return s
}

func (s *BackflowFeatureConsistencyCheckJobDataResponse) Validate() error {
	return dara.Validate(s)
}
