// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureConsistencyCheckJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFeatureConsistencyCheckJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFeatureConsistencyCheckJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetFeatureConsistencyCheckJobConfigResponseBody) *GetFeatureConsistencyCheckJobConfigResponse
	GetBody() *GetFeatureConsistencyCheckJobConfigResponseBody
}

type GetFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureConsistencyCheckJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) GetBody() *GetFeatureConsistencyCheckJobConfigResponseBody {
	return s.Body
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *GetFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *GetFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) SetBody(v *GetFeatureConsistencyCheckJobConfigResponseBody) *GetFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

func (s *GetFeatureConsistencyCheckJobConfigResponse) Validate() error {
	return dara.Validate(s)
}
