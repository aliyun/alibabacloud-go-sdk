// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCronExecTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCronExecTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCronExecTimeResponse
	GetStatusCode() *int32
	SetBody(v *GetCronExecTimeResponseBody) *GetCronExecTimeResponse
	GetBody() *GetCronExecTimeResponseBody
}

type GetCronExecTimeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCronExecTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCronExecTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCronExecTimeResponse) GoString() string {
	return s.String()
}

func (s *GetCronExecTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCronExecTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCronExecTimeResponse) GetBody() *GetCronExecTimeResponseBody {
	return s.Body
}

func (s *GetCronExecTimeResponse) SetHeaders(v map[string]*string) *GetCronExecTimeResponse {
	s.Headers = v
	return s
}

func (s *GetCronExecTimeResponse) SetStatusCode(v int32) *GetCronExecTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCronExecTimeResponse) SetBody(v *GetCronExecTimeResponseBody) *GetCronExecTimeResponse {
	s.Body = v
	return s
}

func (s *GetCronExecTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
