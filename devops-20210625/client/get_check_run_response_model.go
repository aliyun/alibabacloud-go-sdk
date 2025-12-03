// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckRunResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckRunResponseBody) *GetCheckRunResponse
	GetBody() *GetCheckRunResponseBody
}

type GetCheckRunResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunResponse) GoString() string {
	return s.String()
}

func (s *GetCheckRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckRunResponse) GetBody() *GetCheckRunResponseBody {
	return s.Body
}

func (s *GetCheckRunResponse) SetHeaders(v map[string]*string) *GetCheckRunResponse {
	s.Headers = v
	return s
}

func (s *GetCheckRunResponse) SetStatusCode(v int32) *GetCheckRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckRunResponse) SetBody(v *GetCheckRunResponseBody) *GetCheckRunResponse {
	s.Body = v
	return s
}

func (s *GetCheckRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
