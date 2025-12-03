// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemActivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemActivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemActivityResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemActivityResponseBody) *GetWorkItemActivityResponse
	GetBody() *GetWorkItemActivityResponseBody
}

type GetWorkItemActivityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemActivityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemActivityResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemActivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemActivityResponse) GetBody() *GetWorkItemActivityResponseBody {
	return s.Body
}

func (s *GetWorkItemActivityResponse) SetHeaders(v map[string]*string) *GetWorkItemActivityResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemActivityResponse) SetStatusCode(v int32) *GetWorkItemActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemActivityResponse) SetBody(v *GetWorkItemActivityResponseBody) *GetWorkItemActivityResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemActivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
