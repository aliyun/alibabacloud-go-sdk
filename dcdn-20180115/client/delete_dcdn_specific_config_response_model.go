// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnSpecificConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnSpecificConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnSpecificConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnSpecificConfigResponseBody) *DeleteDcdnSpecificConfigResponse
	GetBody() *DeleteDcdnSpecificConfigResponseBody
}

type DeleteDcdnSpecificConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnSpecificConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnSpecificConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnSpecificConfigResponse) GetBody() *DeleteDcdnSpecificConfigResponseBody {
	return s.Body
}

func (s *DeleteDcdnSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnSpecificConfigResponse) SetStatusCode(v int32) *DeleteDcdnSpecificConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnSpecificConfigResponse) SetBody(v *DeleteDcdnSpecificConfigResponseBody) *DeleteDcdnSpecificConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnSpecificConfigResponse) Validate() error {
	return dara.Validate(s)
}
