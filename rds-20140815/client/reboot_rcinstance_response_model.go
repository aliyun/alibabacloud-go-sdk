// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebootRCInstanceResponseBody) *RebootRCInstanceResponse
	GetBody() *RebootRCInstanceResponseBody
}

type RebootRCInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootRCInstanceResponse) GetBody() *RebootRCInstanceResponseBody {
	return s.Body
}

func (s *RebootRCInstanceResponse) SetHeaders(v map[string]*string) *RebootRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootRCInstanceResponse) SetStatusCode(v int32) *RebootRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootRCInstanceResponse) SetBody(v *RebootRCInstanceResponseBody) *RebootRCInstanceResponse {
	s.Body = v
	return s
}

func (s *RebootRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
