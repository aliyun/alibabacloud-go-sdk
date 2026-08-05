// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfflineTaskLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfflineTaskLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfflineTaskLogResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfflineTaskLogResponseBody) *ModifyOfflineTaskLogResponse
	GetBody() *ModifyOfflineTaskLogResponseBody
}

type ModifyOfflineTaskLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfflineTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfflineTaskLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfflineTaskLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfflineTaskLogResponse) GetBody() *ModifyOfflineTaskLogResponseBody {
	return s.Body
}

func (s *ModifyOfflineTaskLogResponse) SetHeaders(v map[string]*string) *ModifyOfflineTaskLogResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfflineTaskLogResponse) SetStatusCode(v int32) *ModifyOfflineTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfflineTaskLogResponse) SetBody(v *ModifyOfflineTaskLogResponseBody) *ModifyOfflineTaskLogResponse {
	s.Body = v
	return s
}

func (s *ModifyOfflineTaskLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
