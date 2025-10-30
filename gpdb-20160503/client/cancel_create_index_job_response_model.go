// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateIndexJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCreateIndexJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCreateIndexJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelCreateIndexJobResponseBody) *CancelCreateIndexJobResponse
	GetBody() *CancelCreateIndexJobResponseBody
}

type CancelCreateIndexJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCreateIndexJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCreateIndexJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateIndexJobResponse) GoString() string {
	return s.String()
}

func (s *CancelCreateIndexJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCreateIndexJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCreateIndexJobResponse) GetBody() *CancelCreateIndexJobResponseBody {
	return s.Body
}

func (s *CancelCreateIndexJobResponse) SetHeaders(v map[string]*string) *CancelCreateIndexJobResponse {
	s.Headers = v
	return s
}

func (s *CancelCreateIndexJobResponse) SetStatusCode(v int32) *CancelCreateIndexJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCreateIndexJobResponse) SetBody(v *CancelCreateIndexJobResponseBody) *CancelCreateIndexJobResponse {
	s.Body = v
	return s
}

func (s *CancelCreateIndexJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
