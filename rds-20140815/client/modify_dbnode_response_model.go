// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeResponseBody) *ModifyDBNodeResponse
	GetBody() *ModifyDBNodeResponseBody
}

type ModifyDBNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeResponse) GetBody() *ModifyDBNodeResponseBody {
	return s.Body
}

func (s *ModifyDBNodeResponse) SetHeaders(v map[string]*string) *ModifyDBNodeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeResponse) SetStatusCode(v int32) *ModifyDBNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeResponse) SetBody(v *ModifyDBNodeResponseBody) *ModifyDBNodeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
