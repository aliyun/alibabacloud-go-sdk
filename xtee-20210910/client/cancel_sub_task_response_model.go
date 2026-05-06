// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelSubTaskResponseBody) *CancelSubTaskResponse
	GetBody() *CancelSubTaskResponseBody
}

type CancelSubTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSubTaskResponse) GetBody() *CancelSubTaskResponseBody {
	return s.Body
}

func (s *CancelSubTaskResponse) SetHeaders(v map[string]*string) *CancelSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelSubTaskResponse) SetStatusCode(v int32) *CancelSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSubTaskResponse) SetBody(v *CancelSubTaskResponseBody) *CancelSubTaskResponse {
	s.Body = v
	return s
}

func (s *CancelSubTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
