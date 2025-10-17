// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApsJobResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApsJobResponseBody) *ModifyApsJobResponse
	GetBody() *ModifyApsJobResponseBody
}

type ModifyApsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyApsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApsJobResponse) GetBody() *ModifyApsJobResponseBody {
	return s.Body
}

func (s *ModifyApsJobResponse) SetHeaders(v map[string]*string) *ModifyApsJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyApsJobResponse) SetStatusCode(v int32) *ModifyApsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApsJobResponse) SetBody(v *ModifyApsJobResponseBody) *ModifyApsJobResponse {
	s.Body = v
	return s
}

func (s *ModifyApsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
