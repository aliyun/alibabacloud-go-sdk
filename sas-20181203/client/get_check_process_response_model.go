// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckProcessResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckProcessResponseBody) *GetCheckProcessResponse
	GetBody() *GetCheckProcessResponseBody
}

type GetCheckProcessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckProcessResponse) GoString() string {
	return s.String()
}

func (s *GetCheckProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckProcessResponse) GetBody() *GetCheckProcessResponseBody {
	return s.Body
}

func (s *GetCheckProcessResponse) SetHeaders(v map[string]*string) *GetCheckProcessResponse {
	s.Headers = v
	return s
}

func (s *GetCheckProcessResponse) SetStatusCode(v int32) *GetCheckProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckProcessResponse) SetBody(v *GetCheckProcessResponseBody) *GetCheckProcessResponse {
	s.Body = v
	return s
}

func (s *GetCheckProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
