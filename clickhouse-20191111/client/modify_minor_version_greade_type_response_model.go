// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMinorVersionGreadeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMinorVersionGreadeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMinorVersionGreadeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMinorVersionGreadeTypeResponseBody) *ModifyMinorVersionGreadeTypeResponse
	GetBody() *ModifyMinorVersionGreadeTypeResponseBody
}

type ModifyMinorVersionGreadeTypeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMinorVersionGreadeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMinorVersionGreadeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMinorVersionGreadeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyMinorVersionGreadeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMinorVersionGreadeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMinorVersionGreadeTypeResponse) GetBody() *ModifyMinorVersionGreadeTypeResponseBody {
	return s.Body
}

func (s *ModifyMinorVersionGreadeTypeResponse) SetHeaders(v map[string]*string) *ModifyMinorVersionGreadeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyMinorVersionGreadeTypeResponse) SetStatusCode(v int32) *ModifyMinorVersionGreadeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeResponse) SetBody(v *ModifyMinorVersionGreadeTypeResponseBody) *ModifyMinorVersionGreadeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyMinorVersionGreadeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
