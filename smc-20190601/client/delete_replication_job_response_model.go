// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReplicationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReplicationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReplicationJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReplicationJobResponseBody) *DeleteReplicationJobResponse
	GetBody() *DeleteReplicationJobResponseBody
}

type DeleteReplicationJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReplicationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReplicationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReplicationJobResponse) GetBody() *DeleteReplicationJobResponseBody {
	return s.Body
}

func (s *DeleteReplicationJobResponse) SetHeaders(v map[string]*string) *DeleteReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteReplicationJobResponse) SetStatusCode(v int32) *DeleteReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReplicationJobResponse) SetBody(v *DeleteReplicationJobResponseBody) *DeleteReplicationJobResponse {
	s.Body = v
	return s
}

func (s *DeleteReplicationJobResponse) Validate() error {
	return dara.Validate(s)
}
