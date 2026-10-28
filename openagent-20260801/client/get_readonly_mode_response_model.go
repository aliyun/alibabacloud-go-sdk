// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReadonlyModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReadonlyModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReadonlyModeResponse
	GetStatusCode() *int32
	SetBody(v *GetReadonlyModeResponseBody) *GetReadonlyModeResponse
	GetBody() *GetReadonlyModeResponseBody
}

type GetReadonlyModeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReadonlyModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReadonlyModeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReadonlyModeResponse) GoString() string {
	return s.String()
}

func (s *GetReadonlyModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReadonlyModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReadonlyModeResponse) GetBody() *GetReadonlyModeResponseBody {
	return s.Body
}

func (s *GetReadonlyModeResponse) SetHeaders(v map[string]*string) *GetReadonlyModeResponse {
	s.Headers = v
	return s
}

func (s *GetReadonlyModeResponse) SetStatusCode(v int32) *GetReadonlyModeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReadonlyModeResponse) SetBody(v *GetReadonlyModeResponseBody) *GetReadonlyModeResponse {
	s.Body = v
	return s
}

func (s *GetReadonlyModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
