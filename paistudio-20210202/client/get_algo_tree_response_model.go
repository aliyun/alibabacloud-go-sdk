// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgoTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlgoTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlgoTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetAlgoTreeResponseBody) *GetAlgoTreeResponse
	GetBody() *GetAlgoTreeResponseBody
}

type GetAlgoTreeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgoTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgoTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlgoTreeResponse) GoString() string {
	return s.String()
}

func (s *GetAlgoTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlgoTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlgoTreeResponse) GetBody() *GetAlgoTreeResponseBody {
	return s.Body
}

func (s *GetAlgoTreeResponse) SetHeaders(v map[string]*string) *GetAlgoTreeResponse {
	s.Headers = v
	return s
}

func (s *GetAlgoTreeResponse) SetStatusCode(v int32) *GetAlgoTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgoTreeResponse) SetBody(v *GetAlgoTreeResponseBody) *GetAlgoTreeResponse {
	s.Body = v
	return s
}

func (s *GetAlgoTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
