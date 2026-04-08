// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeResponseBody) *GetNodeResponse
	GetBody() *GetNodeResponseBody
}

type GetNodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeResponse) GetBody() *GetNodeResponseBody {
	return s.Body
}

func (s *GetNodeResponse) SetHeaders(v map[string]*string) *GetNodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeResponse) SetStatusCode(v int32) *GetNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeResponse) SetBody(v *GetNodeResponseBody) *GetNodeResponse {
	s.Body = v
	return s
}

func (s *GetNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
