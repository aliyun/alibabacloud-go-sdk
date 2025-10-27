// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefenceCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefenceCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefenceCountResponse
	GetStatusCode() *int32
	SetBody(v *GetDefenceCountResponseBody) *GetDefenceCountResponse
	GetBody() *GetDefenceCountResponseBody
}

type GetDefenceCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefenceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefenceCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefenceCountResponse) GoString() string {
	return s.String()
}

func (s *GetDefenceCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefenceCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefenceCountResponse) GetBody() *GetDefenceCountResponseBody {
	return s.Body
}

func (s *GetDefenceCountResponse) SetHeaders(v map[string]*string) *GetDefenceCountResponse {
	s.Headers = v
	return s
}

func (s *GetDefenceCountResponse) SetStatusCode(v int32) *GetDefenceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefenceCountResponse) SetBody(v *GetDefenceCountResponseBody) *GetDefenceCountResponse {
	s.Body = v
	return s
}

func (s *GetDefenceCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
