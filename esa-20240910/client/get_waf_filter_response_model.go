// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWafFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWafFilterResponse
	GetStatusCode() *int32
	SetBody(v *GetWafFilterResponseBody) *GetWafFilterResponse
	GetBody() *GetWafFilterResponseBody
}

type GetWafFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponse) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWafFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWafFilterResponse) GetBody() *GetWafFilterResponseBody {
	return s.Body
}

func (s *GetWafFilterResponse) SetHeaders(v map[string]*string) *GetWafFilterResponse {
	s.Headers = v
	return s
}

func (s *GetWafFilterResponse) SetStatusCode(v int32) *GetWafFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafFilterResponse) SetBody(v *GetWafFilterResponseBody) *GetWafFilterResponse {
	s.Body = v
	return s
}

func (s *GetWafFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
