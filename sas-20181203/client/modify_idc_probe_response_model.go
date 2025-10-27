// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIdcProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIdcProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIdcProbeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIdcProbeResponseBody) *ModifyIdcProbeResponse
	GetBody() *ModifyIdcProbeResponseBody
}

type ModifyIdcProbeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIdcProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIdcProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIdcProbeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIdcProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIdcProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIdcProbeResponse) GetBody() *ModifyIdcProbeResponseBody {
	return s.Body
}

func (s *ModifyIdcProbeResponse) SetHeaders(v map[string]*string) *ModifyIdcProbeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIdcProbeResponse) SetStatusCode(v int32) *ModifyIdcProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIdcProbeResponse) SetBody(v *ModifyIdcProbeResponseBody) *ModifyIdcProbeResponse {
	s.Body = v
	return s
}

func (s *ModifyIdcProbeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
