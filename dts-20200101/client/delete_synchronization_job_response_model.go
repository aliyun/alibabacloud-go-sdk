// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSynchronizationJobResponseBody) *DeleteSynchronizationJobResponse
	GetBody() *DeleteSynchronizationJobResponseBody
}

type DeleteSynchronizationJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSynchronizationJobResponse) GetBody() *DeleteSynchronizationJobResponseBody {
	return s.Body
}

func (s *DeleteSynchronizationJobResponse) SetHeaders(v map[string]*string) *DeleteSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSynchronizationJobResponse) SetStatusCode(v int32) *DeleteSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSynchronizationJobResponse) SetBody(v *DeleteSynchronizationJobResponseBody) *DeleteSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *DeleteSynchronizationJobResponse) Validate() error {
	return dara.Validate(s)
}
