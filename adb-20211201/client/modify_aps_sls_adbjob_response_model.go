// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsSlsADBJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApsSlsADBJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApsSlsADBJobResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApsSlsADBJobResponseBody) *ModifyApsSlsADBJobResponse
	GetBody() *ModifyApsSlsADBJobResponseBody
}

type ModifyApsSlsADBJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApsSlsADBJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApsSlsADBJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsSlsADBJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyApsSlsADBJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApsSlsADBJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApsSlsADBJobResponse) GetBody() *ModifyApsSlsADBJobResponseBody {
	return s.Body
}

func (s *ModifyApsSlsADBJobResponse) SetHeaders(v map[string]*string) *ModifyApsSlsADBJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyApsSlsADBJobResponse) SetStatusCode(v int32) *ModifyApsSlsADBJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApsSlsADBJobResponse) SetBody(v *ModifyApsSlsADBJobResponseBody) *ModifyApsSlsADBJobResponse {
	s.Body = v
	return s
}

func (s *ModifyApsSlsADBJobResponse) Validate() error {
	return dara.Validate(s)
}
