// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgOneKeyDeleteTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgOneKeyDeleteTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgOneKeyDeleteTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAgOneKeyDeleteTaskResponseBody) *GetAgOneKeyDeleteTaskResponse
	GetBody() *GetAgOneKeyDeleteTaskResponseBody
}

type GetAgOneKeyDeleteTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgOneKeyDeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgOneKeyDeleteTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgOneKeyDeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAgOneKeyDeleteTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgOneKeyDeleteTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgOneKeyDeleteTaskResponse) GetBody() *GetAgOneKeyDeleteTaskResponseBody {
	return s.Body
}

func (s *GetAgOneKeyDeleteTaskResponse) SetHeaders(v map[string]*string) *GetAgOneKeyDeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponse) SetStatusCode(v int32) *GetAgOneKeyDeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponse) SetBody(v *GetAgOneKeyDeleteTaskResponseBody) *GetAgOneKeyDeleteTaskResponse {
	s.Body = v
	return s
}

func (s *GetAgOneKeyDeleteTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
