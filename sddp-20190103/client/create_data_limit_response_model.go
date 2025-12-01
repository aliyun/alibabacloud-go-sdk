// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataLimitResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataLimitResponseBody) *CreateDataLimitResponse
	GetBody() *CreateDataLimitResponseBody
}

type CreateDataLimitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLimitResponse) GoString() string {
	return s.String()
}

func (s *CreateDataLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataLimitResponse) GetBody() *CreateDataLimitResponseBody {
	return s.Body
}

func (s *CreateDataLimitResponse) SetHeaders(v map[string]*string) *CreateDataLimitResponse {
	s.Headers = v
	return s
}

func (s *CreateDataLimitResponse) SetStatusCode(v int32) *CreateDataLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataLimitResponse) SetBody(v *CreateDataLimitResponseBody) *CreateDataLimitResponse {
	s.Body = v
	return s
}

func (s *CreateDataLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
