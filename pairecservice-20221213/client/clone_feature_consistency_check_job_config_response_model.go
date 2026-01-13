// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneFeatureConsistencyCheckJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneFeatureConsistencyCheckJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneFeatureConsistencyCheckJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *CloneFeatureConsistencyCheckJobConfigResponseBody) *CloneFeatureConsistencyCheckJobConfigResponse
	GetBody() *CloneFeatureConsistencyCheckJobConfigResponseBody
}

type CloneFeatureConsistencyCheckJobConfigResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneFeatureConsistencyCheckJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneFeatureConsistencyCheckJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneFeatureConsistencyCheckJobConfigResponse) GoString() string {
	return s.String()
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) GetBody() *CloneFeatureConsistencyCheckJobConfigResponseBody {
	return s.Body
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) SetHeaders(v map[string]*string) *CloneFeatureConsistencyCheckJobConfigResponse {
	s.Headers = v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) SetStatusCode(v int32) *CloneFeatureConsistencyCheckJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) SetBody(v *CloneFeatureConsistencyCheckJobConfigResponseBody) *CloneFeatureConsistencyCheckJobConfigResponse {
	s.Body = v
	return s
}

func (s *CloneFeatureConsistencyCheckJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
