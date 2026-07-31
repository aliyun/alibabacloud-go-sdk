// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSemanticViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSemanticViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSemanticViewResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSemanticViewResponseBody) *DeleteSemanticViewResponse
	GetBody() *DeleteSemanticViewResponseBody
}

type DeleteSemanticViewResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSemanticViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSemanticViewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSemanticViewResponse) GoString() string {
	return s.String()
}

func (s *DeleteSemanticViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSemanticViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSemanticViewResponse) GetBody() *DeleteSemanticViewResponseBody {
	return s.Body
}

func (s *DeleteSemanticViewResponse) SetHeaders(v map[string]*string) *DeleteSemanticViewResponse {
	s.Headers = v
	return s
}

func (s *DeleteSemanticViewResponse) SetStatusCode(v int32) *DeleteSemanticViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSemanticViewResponse) SetBody(v *DeleteSemanticViewResponseBody) *DeleteSemanticViewResponse {
	s.Body = v
	return s
}

func (s *DeleteSemanticViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
