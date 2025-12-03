// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupByPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupByPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupByPathResponse
	GetStatusCode() *int32
	SetBody(v *GetGroupByPathResponseBody) *GetGroupByPathResponse
	GetBody() *GetGroupByPathResponseBody
}

type GetGroupByPathResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupByPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupByPathResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupByPathResponse) GoString() string {
	return s.String()
}

func (s *GetGroupByPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupByPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupByPathResponse) GetBody() *GetGroupByPathResponseBody {
	return s.Body
}

func (s *GetGroupByPathResponse) SetHeaders(v map[string]*string) *GetGroupByPathResponse {
	s.Headers = v
	return s
}

func (s *GetGroupByPathResponse) SetStatusCode(v int32) *GetGroupByPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupByPathResponse) SetBody(v *GetGroupByPathResponseBody) *GetGroupByPathResponse {
	s.Body = v
	return s
}

func (s *GetGroupByPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
