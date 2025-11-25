// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAuthRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceAuthRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceAuthRangeResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceAuthRangeResponseBody) *GetInstanceAuthRangeResponse
	GetBody() *GetInstanceAuthRangeResponseBody
}

type GetInstanceAuthRangeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAuthRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAuthRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAuthRangeResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAuthRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceAuthRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceAuthRangeResponse) GetBody() *GetInstanceAuthRangeResponseBody {
	return s.Body
}

func (s *GetInstanceAuthRangeResponse) SetHeaders(v map[string]*string) *GetInstanceAuthRangeResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAuthRangeResponse) SetStatusCode(v int32) *GetInstanceAuthRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAuthRangeResponse) SetBody(v *GetInstanceAuthRangeResponseBody) *GetInstanceAuthRangeResponse {
	s.Body = v
	return s
}

func (s *GetInstanceAuthRangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
