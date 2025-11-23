// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDBTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableDBTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableDBTopologyResponse
	GetStatusCode() *int32
	SetBody(v *GetTableDBTopologyResponseBody) *GetTableDBTopologyResponse
	GetBody() *GetTableDBTopologyResponseBody
}

type GetTableDBTopologyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableDBTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableDBTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableDBTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableDBTopologyResponse) GetBody() *GetTableDBTopologyResponseBody {
	return s.Body
}

func (s *GetTableDBTopologyResponse) SetHeaders(v map[string]*string) *GetTableDBTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetTableDBTopologyResponse) SetStatusCode(v int32) *GetTableDBTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableDBTopologyResponse) SetBody(v *GetTableDBTopologyResponseBody) *GetTableDBTopologyResponse {
	s.Body = v
	return s
}

func (s *GetTableDBTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
