// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourcePackStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSourcePackStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSourcePackStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSourcePackStatusResponseBody) *GetSourcePackStatusResponse
	GetBody() *GetSourcePackStatusResponseBody
}

type GetSourcePackStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSourcePackStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSourcePackStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSourcePackStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSourcePackStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSourcePackStatusResponse) GetBody() *GetSourcePackStatusResponseBody {
	return s.Body
}

func (s *GetSourcePackStatusResponse) SetHeaders(v map[string]*string) *GetSourcePackStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSourcePackStatusResponse) SetStatusCode(v int32) *GetSourcePackStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourcePackStatusResponse) SetBody(v *GetSourcePackStatusResponseBody) *GetSourcePackStatusResponse {
	s.Body = v
	return s
}

func (s *GetSourcePackStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
