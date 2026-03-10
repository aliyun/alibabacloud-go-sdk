// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighlightTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHighlightTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHighlightTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateHighlightTaskResponseBody) *CreateHighlightTaskResponse
	GetBody() *CreateHighlightTaskResponseBody
}

type CreateHighlightTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHighlightTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHighlightTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHighlightTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHighlightTaskResponse) GetBody() *CreateHighlightTaskResponseBody {
	return s.Body
}

func (s *CreateHighlightTaskResponse) SetHeaders(v map[string]*string) *CreateHighlightTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateHighlightTaskResponse) SetStatusCode(v int32) *CreateHighlightTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHighlightTaskResponse) SetBody(v *CreateHighlightTaskResponseBody) *CreateHighlightTaskResponse {
	s.Body = v
	return s
}

func (s *CreateHighlightTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
