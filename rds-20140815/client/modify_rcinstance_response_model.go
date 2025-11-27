// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCInstanceResponseBody) *ModifyRCInstanceResponse
	GetBody() *ModifyRCInstanceResponseBody
}

type ModifyRCInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCInstanceResponse) GetBody() *ModifyRCInstanceResponseBody {
	return s.Body
}

func (s *ModifyRCInstanceResponse) SetHeaders(v map[string]*string) *ModifyRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCInstanceResponse) SetStatusCode(v int32) *ModifyRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCInstanceResponse) SetBody(v *ModifyRCInstanceResponseBody) *ModifyRCInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
