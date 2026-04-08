// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDIJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDIJobResponseBody) *DeleteDIJobResponse
	GetBody() *DeleteDIJobResponseBody
}

type DeleteDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDIJobResponse) GetBody() *DeleteDIJobResponseBody {
	return s.Body
}

func (s *DeleteDIJobResponse) SetHeaders(v map[string]*string) *DeleteDIJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDIJobResponse) SetStatusCode(v int32) *DeleteDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDIJobResponse) SetBody(v *DeleteDIJobResponseBody) *DeleteDIJobResponse {
	s.Body = v
	return s
}

func (s *DeleteDIJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
