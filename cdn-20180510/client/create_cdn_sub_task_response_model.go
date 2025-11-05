// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCdnSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCdnSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateCdnSubTaskResponseBody) *CreateCdnSubTaskResponse
	GetBody() *CreateCdnSubTaskResponseBody
}

type CreateCdnSubTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdnSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCdnSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCdnSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCdnSubTaskResponse) GetBody() *CreateCdnSubTaskResponseBody {
	return s.Body
}

func (s *CreateCdnSubTaskResponse) SetHeaders(v map[string]*string) *CreateCdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCdnSubTaskResponse) SetStatusCode(v int32) *CreateCdnSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdnSubTaskResponse) SetBody(v *CreateCdnSubTaskResponseBody) *CreateCdnSubTaskResponse {
	s.Body = v
	return s
}

func (s *CreateCdnSubTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
