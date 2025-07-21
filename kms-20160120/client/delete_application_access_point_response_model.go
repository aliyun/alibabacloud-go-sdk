// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationAccessPointResponseBody) *DeleteApplicationAccessPointResponse
	GetBody() *DeleteApplicationAccessPointResponseBody
}

type DeleteApplicationAccessPointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationAccessPointResponse) GetBody() *DeleteApplicationAccessPointResponseBody {
	return s.Body
}

func (s *DeleteApplicationAccessPointResponse) SetHeaders(v map[string]*string) *DeleteApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationAccessPointResponse) SetStatusCode(v int32) *DeleteApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationAccessPointResponse) SetBody(v *DeleteApplicationAccessPointResponseBody) *DeleteApplicationAccessPointResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationAccessPointResponse) Validate() error {
	return dara.Validate(s)
}
