// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IncreaseNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IncreaseNodesResponse
	GetStatusCode() *int32
	SetBody(v *IncreaseNodesResponseBody) *IncreaseNodesResponse
	GetBody() *IncreaseNodesResponseBody
}

type IncreaseNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IncreaseNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IncreaseNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s IncreaseNodesResponse) GoString() string {
	return s.String()
}

func (s *IncreaseNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IncreaseNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IncreaseNodesResponse) GetBody() *IncreaseNodesResponseBody {
	return s.Body
}

func (s *IncreaseNodesResponse) SetHeaders(v map[string]*string) *IncreaseNodesResponse {
	s.Headers = v
	return s
}

func (s *IncreaseNodesResponse) SetStatusCode(v int32) *IncreaseNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseNodesResponse) SetBody(v *IncreaseNodesResponseBody) *IncreaseNodesResponse {
	s.Body = v
	return s
}

func (s *IncreaseNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
