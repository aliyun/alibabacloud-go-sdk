// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableTopologyResponse
	GetStatusCode() *int32
	SetBody(v *GetTableTopologyResponseBody) *GetTableTopologyResponse
	GetBody() *GetTableTopologyResponseBody
}

type GetTableTopologyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableTopologyResponse) GetBody() *GetTableTopologyResponseBody {
	return s.Body
}

func (s *GetTableTopologyResponse) SetHeaders(v map[string]*string) *GetTableTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetTableTopologyResponse) SetStatusCode(v int32) *GetTableTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableTopologyResponse) SetBody(v *GetTableTopologyResponseBody) *GetTableTopologyResponse {
	s.Body = v
	return s
}

func (s *GetTableTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
