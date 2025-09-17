// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReplicationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartReplicationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartReplicationJobResponse
	GetStatusCode() *int32
	SetBody(v *StartReplicationJobResponseBody) *StartReplicationJobResponse
	GetBody() *StartReplicationJobResponseBody
}

type StartReplicationJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartReplicationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *StartReplicationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartReplicationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartReplicationJobResponse) GetBody() *StartReplicationJobResponseBody {
	return s.Body
}

func (s *StartReplicationJobResponse) SetHeaders(v map[string]*string) *StartReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *StartReplicationJobResponse) SetStatusCode(v int32) *StartReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartReplicationJobResponse) SetBody(v *StartReplicationJobResponseBody) *StartReplicationJobResponse {
	s.Body = v
	return s
}

func (s *StartReplicationJobResponse) Validate() error {
	return dara.Validate(s)
}
