// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSpecificStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveSpecificStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveSpecificStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveSpecificStagingConfigResponseBody) *DeleteLiveSpecificStagingConfigResponse
	GetBody() *DeleteLiveSpecificStagingConfigResponseBody
}

type DeleteLiveSpecificStagingConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveSpecificStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveSpecificStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSpecificStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveSpecificStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveSpecificStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveSpecificStagingConfigResponse) GetBody() *DeleteLiveSpecificStagingConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveSpecificStagingConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveSpecificStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveSpecificStagingConfigResponse) SetStatusCode(v int32) *DeleteLiveSpecificStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveSpecificStagingConfigResponse) SetBody(v *DeleteLiveSpecificStagingConfigResponseBody) *DeleteLiveSpecificStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveSpecificStagingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
