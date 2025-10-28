// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetModelVersionResponseBody) *GetModelVersionResponse
	GetBody() *GetModelVersionResponseBody
}

type GetModelVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelVersionResponse) GoString() string {
	return s.String()
}

func (s *GetModelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelVersionResponse) GetBody() *GetModelVersionResponseBody {
	return s.Body
}

func (s *GetModelVersionResponse) SetHeaders(v map[string]*string) *GetModelVersionResponse {
	s.Headers = v
	return s
}

func (s *GetModelVersionResponse) SetStatusCode(v int32) *GetModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelVersionResponse) SetBody(v *GetModelVersionResponseBody) *GetModelVersionResponse {
	s.Body = v
	return s
}

func (s *GetModelVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
