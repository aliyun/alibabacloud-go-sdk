// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAINodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAINodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAINodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAINodeResponseBody) *DeleteAINodeResponse
	GetBody() *DeleteAINodeResponseBody
}

type DeleteAINodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAINodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAINodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAINodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteAINodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAINodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAINodeResponse) GetBody() *DeleteAINodeResponseBody {
	return s.Body
}

func (s *DeleteAINodeResponse) SetHeaders(v map[string]*string) *DeleteAINodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteAINodeResponse) SetStatusCode(v int32) *DeleteAINodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAINodeResponse) SetBody(v *DeleteAINodeResponseBody) *DeleteAINodeResponse {
	s.Body = v
	return s
}

func (s *DeleteAINodeResponse) Validate() error {
	return dara.Validate(s)
}
