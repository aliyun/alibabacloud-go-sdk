// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagHaResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagHaResponseBody) *ModifySagHaResponse
	GetBody() *ModifySagHaResponseBody
}

type ModifySagHaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagHaResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagHaResponse) GoString() string {
	return s.String()
}

func (s *ModifySagHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagHaResponse) GetBody() *ModifySagHaResponseBody {
	return s.Body
}

func (s *ModifySagHaResponse) SetHeaders(v map[string]*string) *ModifySagHaResponse {
	s.Headers = v
	return s
}

func (s *ModifySagHaResponse) SetStatusCode(v int32) *ModifySagHaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagHaResponse) SetBody(v *ModifySagHaResponseBody) *ModifySagHaResponse {
	s.Body = v
	return s
}

func (s *ModifySagHaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
