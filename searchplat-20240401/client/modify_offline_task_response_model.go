// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOfflineTaskResponseBody) *ModifyOfflineTaskResponse
	GetBody() *ModifyOfflineTaskResponseBody
}

type ModifyOfflineTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOfflineTaskResponse) GetBody() *ModifyOfflineTaskResponseBody {
	return s.Body
}

func (s *ModifyOfflineTaskResponse) SetHeaders(v map[string]*string) *ModifyOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfflineTaskResponse) SetStatusCode(v int32) *ModifyOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOfflineTaskResponse) SetBody(v *ModifyOfflineTaskResponseBody) *ModifyOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
