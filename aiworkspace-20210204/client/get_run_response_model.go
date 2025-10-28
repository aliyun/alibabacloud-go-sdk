// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRunResponse
	GetStatusCode() *int32
	SetBody(v *Run) *GetRunResponse
	GetBody() *Run
}

type GetRunResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Run               `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRunResponse) GoString() string {
	return s.String()
}

func (s *GetRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRunResponse) GetBody() *Run {
	return s.Body
}

func (s *GetRunResponse) SetHeaders(v map[string]*string) *GetRunResponse {
	s.Headers = v
	return s
}

func (s *GetRunResponse) SetStatusCode(v int32) *GetRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunResponse) SetBody(v *Run) *GetRunResponse {
	s.Body = v
	return s
}

func (s *GetRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
