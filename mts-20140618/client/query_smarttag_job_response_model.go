// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmarttagJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmarttagJobResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmarttagJobResponseBody) *QuerySmarttagJobResponse
	GetBody() *QuerySmarttagJobResponseBody
}

type QuerySmarttagJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmarttagJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmarttagJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagJobResponse) GoString() string {
	return s.String()
}

func (s *QuerySmarttagJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmarttagJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmarttagJobResponse) GetBody() *QuerySmarttagJobResponseBody {
	return s.Body
}

func (s *QuerySmarttagJobResponse) SetHeaders(v map[string]*string) *QuerySmarttagJobResponse {
	s.Headers = v
	return s
}

func (s *QuerySmarttagJobResponse) SetStatusCode(v int32) *QuerySmarttagJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmarttagJobResponse) SetBody(v *QuerySmarttagJobResponseBody) *QuerySmarttagJobResponse {
	s.Body = v
	return s
}

func (s *QuerySmarttagJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
