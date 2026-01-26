// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumExceptionStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumExceptionStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumExceptionStackResponse
	GetStatusCode() *int32
	SetBody(v *GetRumExceptionStackResponseBody) *GetRumExceptionStackResponse
	GetBody() *GetRumExceptionStackResponseBody
}

type GetRumExceptionStackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumExceptionStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumExceptionStackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumExceptionStackResponse) GoString() string {
	return s.String()
}

func (s *GetRumExceptionStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumExceptionStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumExceptionStackResponse) GetBody() *GetRumExceptionStackResponseBody {
	return s.Body
}

func (s *GetRumExceptionStackResponse) SetHeaders(v map[string]*string) *GetRumExceptionStackResponse {
	s.Headers = v
	return s
}

func (s *GetRumExceptionStackResponse) SetStatusCode(v int32) *GetRumExceptionStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumExceptionStackResponse) SetBody(v *GetRumExceptionStackResponseBody) *GetRumExceptionStackResponse {
	s.Body = v
	return s
}

func (s *GetRumExceptionStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
