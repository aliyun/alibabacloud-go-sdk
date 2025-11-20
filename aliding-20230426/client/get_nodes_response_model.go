// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodesResponse
	GetStatusCode() *int32
	SetBody(v *GetNodesResponseBody) *GetNodesResponse
	GetBody() *GetNodesResponseBody
}

type GetNodesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodesResponse) GoString() string {
	return s.String()
}

func (s *GetNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodesResponse) GetBody() *GetNodesResponseBody {
	return s.Body
}

func (s *GetNodesResponse) SetHeaders(v map[string]*string) *GetNodesResponse {
	s.Headers = v
	return s
}

func (s *GetNodesResponse) SetStatusCode(v int32) *GetNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodesResponse) SetBody(v *GetNodesResponseBody) *GetNodesResponse {
	s.Body = v
	return s
}

func (s *GetNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
