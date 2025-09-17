// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReplicationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReplicationJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateReplicationJobResponseBody) *CreateReplicationJobResponse
	GetBody() *CreateReplicationJobResponseBody
}

type CreateReplicationJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReplicationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReplicationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReplicationJobResponse) GetBody() *CreateReplicationJobResponseBody {
	return s.Body
}

func (s *CreateReplicationJobResponse) SetHeaders(v map[string]*string) *CreateReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateReplicationJobResponse) SetStatusCode(v int32) *CreateReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReplicationJobResponse) SetBody(v *CreateReplicationJobResponseBody) *CreateReplicationJobResponse {
	s.Body = v
	return s
}

func (s *CreateReplicationJobResponse) Validate() error {
	return dara.Validate(s)
}
