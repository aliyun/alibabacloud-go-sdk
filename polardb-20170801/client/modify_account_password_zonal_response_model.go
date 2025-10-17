// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountPasswordZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountPasswordZonalResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountPasswordZonalResponseBody) *ModifyAccountPasswordZonalResponse
	GetBody() *ModifyAccountPasswordZonalResponseBody
}

type ModifyAccountPasswordZonalResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPasswordZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountPasswordZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordZonalResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountPasswordZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountPasswordZonalResponse) GetBody() *ModifyAccountPasswordZonalResponseBody {
	return s.Body
}

func (s *ModifyAccountPasswordZonalResponse) SetHeaders(v map[string]*string) *ModifyAccountPasswordZonalResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPasswordZonalResponse) SetStatusCode(v int32) *ModifyAccountPasswordZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPasswordZonalResponse) SetBody(v *ModifyAccountPasswordZonalResponseBody) *ModifyAccountPasswordZonalResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountPasswordZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
