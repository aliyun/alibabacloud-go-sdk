// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaIndexJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaIndexJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaIndexJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaIndexJobResponseBody) *QueryMediaIndexJobResponse
	GetBody() *QueryMediaIndexJobResponseBody
}

type QueryMediaIndexJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaIndexJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaIndexJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaIndexJobResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaIndexJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaIndexJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaIndexJobResponse) GetBody() *QueryMediaIndexJobResponseBody {
	return s.Body
}

func (s *QueryMediaIndexJobResponse) SetHeaders(v map[string]*string) *QueryMediaIndexJobResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaIndexJobResponse) SetStatusCode(v int32) *QueryMediaIndexJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaIndexJobResponse) SetBody(v *QueryMediaIndexJobResponseBody) *QueryMediaIndexJobResponse {
	s.Body = v
	return s
}

func (s *QueryMediaIndexJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
