// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetextJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoDetextJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoDetextJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoDetextJobResponseBody) *SubmitVideoDetextJobResponse
	GetBody() *SubmitVideoDetextJobResponseBody
}

type SubmitVideoDetextJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoDetextJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoDetextJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetextJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetextJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoDetextJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoDetextJobResponse) GetBody() *SubmitVideoDetextJobResponseBody {
	return s.Body
}

func (s *SubmitVideoDetextJobResponse) SetHeaders(v map[string]*string) *SubmitVideoDetextJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoDetextJobResponse) SetStatusCode(v int32) *SubmitVideoDetextJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoDetextJobResponse) SetBody(v *SubmitVideoDetextJobResponseBody) *SubmitVideoDetextJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoDetextJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
