// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoginProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoginProfileResponse
	GetStatusCode() *int32
	SetBody(v *GetLoginProfileResponseBody) *GetLoginProfileResponse
	GetBody() *GetLoginProfileResponseBody
}

type GetLoginProfileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoginProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoginProfileResponse) GetBody() *GetLoginProfileResponseBody {
	return s.Body
}

func (s *GetLoginProfileResponse) SetHeaders(v map[string]*string) *GetLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *GetLoginProfileResponse) SetStatusCode(v int32) *GetLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginProfileResponse) SetBody(v *GetLoginProfileResponseBody) *GetLoginProfileResponse {
	s.Body = v
	return s
}

func (s *GetLoginProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
