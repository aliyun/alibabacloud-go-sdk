// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTagOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTagOptionResponse
	GetStatusCode() *int32
	SetBody(v *GetTagOptionResponseBody) *GetTagOptionResponse
	GetBody() *GetTagOptionResponseBody
}

type GetTagOptionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTagOptionResponse) GoString() string {
	return s.String()
}

func (s *GetTagOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTagOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTagOptionResponse) GetBody() *GetTagOptionResponseBody {
	return s.Body
}

func (s *GetTagOptionResponse) SetHeaders(v map[string]*string) *GetTagOptionResponse {
	s.Headers = v
	return s
}

func (s *GetTagOptionResponse) SetStatusCode(v int32) *GetTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagOptionResponse) SetBody(v *GetTagOptionResponseBody) *GetTagOptionResponse {
	s.Body = v
	return s
}

func (s *GetTagOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
