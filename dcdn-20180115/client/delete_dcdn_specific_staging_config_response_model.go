// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSpecificStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnSpecificStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnSpecificStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnSpecificStagingConfigResponseBody) *DeleteDcdnSpecificStagingConfigResponse
	GetBody() *DeleteDcdnSpecificStagingConfigResponseBody
}

type DeleteDcdnSpecificStagingConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnSpecificStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnSpecificStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSpecificStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnSpecificStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnSpecificStagingConfigResponse) GetBody() *DeleteDcdnSpecificStagingConfigResponseBody {
	return s.Body
}

func (s *DeleteDcdnSpecificStagingConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnSpecificStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigResponse) SetStatusCode(v int32) *DeleteDcdnSpecificStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigResponse) SetBody(v *DeleteDcdnSpecificStagingConfigResponseBody) *DeleteDcdnSpecificStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
