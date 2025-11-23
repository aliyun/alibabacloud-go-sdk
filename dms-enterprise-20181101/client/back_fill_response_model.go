// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackFillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BackFillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BackFillResponse
	GetStatusCode() *int32
	SetBody(v *BackFillResponseBody) *BackFillResponse
	GetBody() *BackFillResponseBody
}

type BackFillResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackFillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackFillResponse) String() string {
	return dara.Prettify(s)
}

func (s BackFillResponse) GoString() string {
	return s.String()
}

func (s *BackFillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BackFillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BackFillResponse) GetBody() *BackFillResponseBody {
	return s.Body
}

func (s *BackFillResponse) SetHeaders(v map[string]*string) *BackFillResponse {
	s.Headers = v
	return s
}

func (s *BackFillResponse) SetStatusCode(v int32) *BackFillResponse {
	s.StatusCode = &v
	return s
}

func (s *BackFillResponse) SetBody(v *BackFillResponseBody) *BackFillResponse {
	s.Body = v
	return s
}

func (s *BackFillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
