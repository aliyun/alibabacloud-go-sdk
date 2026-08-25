// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelConnectionResponse
	GetStatusCode() *int32
	SetBody(v *GetModelConnectionResponseBody) *GetModelConnectionResponse
	GetBody() *GetModelConnectionResponseBody
}

type GetModelConnectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelConnectionResponse) GoString() string {
	return s.String()
}

func (s *GetModelConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelConnectionResponse) GetBody() *GetModelConnectionResponseBody {
	return s.Body
}

func (s *GetModelConnectionResponse) SetHeaders(v map[string]*string) *GetModelConnectionResponse {
	s.Headers = v
	return s
}

func (s *GetModelConnectionResponse) SetStatusCode(v int32) *GetModelConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelConnectionResponse) SetBody(v *GetModelConnectionResponseBody) *GetModelConnectionResponse {
	s.Body = v
	return s
}

func (s *GetModelConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
