// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpecificStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSpecificStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSpecificStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSpecificStagingConfigResponseBody) *DeleteSpecificStagingConfigResponse
	GetBody() *DeleteSpecificStagingConfigResponseBody
}

type DeleteSpecificStagingConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSpecificStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSpecificStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpecificStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpecificStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSpecificStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSpecificStagingConfigResponse) GetBody() *DeleteSpecificStagingConfigResponseBody {
	return s.Body
}

func (s *DeleteSpecificStagingConfigResponse) SetHeaders(v map[string]*string) *DeleteSpecificStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpecificStagingConfigResponse) SetStatusCode(v int32) *DeleteSpecificStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSpecificStagingConfigResponse) SetBody(v *DeleteSpecificStagingConfigResponseBody) *DeleteSpecificStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteSpecificStagingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
