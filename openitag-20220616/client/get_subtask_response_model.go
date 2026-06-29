// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubtaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubtaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubtaskResponse
	GetStatusCode() *int32
	SetBody(v *GetSubtaskResponseBody) *GetSubtaskResponse
	GetBody() *GetSubtaskResponseBody
}

type GetSubtaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubtaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubtaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubtaskResponse) GoString() string {
	return s.String()
}

func (s *GetSubtaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubtaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubtaskResponse) GetBody() *GetSubtaskResponseBody {
	return s.Body
}

func (s *GetSubtaskResponse) SetHeaders(v map[string]*string) *GetSubtaskResponse {
	s.Headers = v
	return s
}

func (s *GetSubtaskResponse) SetStatusCode(v int32) *GetSubtaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubtaskResponse) SetBody(v *GetSubtaskResponseBody) *GetSubtaskResponse {
	s.Body = v
	return s
}

func (s *GetSubtaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
