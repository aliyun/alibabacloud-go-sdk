// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDtsJobResponseBody) *ModifyDtsJobResponse
	GetBody() *ModifyDtsJobResponseBody
}

type ModifyDtsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDtsJobResponse) GetBody() *ModifyDtsJobResponseBody {
	return s.Body
}

func (s *ModifyDtsJobResponse) SetHeaders(v map[string]*string) *ModifyDtsJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobResponse) SetStatusCode(v int32) *ModifyDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDtsJobResponse) SetBody(v *ModifyDtsJobResponseBody) *ModifyDtsJobResponse {
	s.Body = v
	return s
}

func (s *ModifyDtsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
