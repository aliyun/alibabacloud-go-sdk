// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserSlsLogStorageTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserSlsLogStorageTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserSlsLogStorageTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserSlsLogStorageTimeResponseBody) *ModifyUserSlsLogStorageTimeResponse
	GetBody() *ModifyUserSlsLogStorageTimeResponseBody
}

type ModifyUserSlsLogStorageTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserSlsLogStorageTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserSlsLogStorageTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserSlsLogStorageTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserSlsLogStorageTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserSlsLogStorageTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserSlsLogStorageTimeResponse) GetBody() *ModifyUserSlsLogStorageTimeResponseBody {
	return s.Body
}

func (s *ModifyUserSlsLogStorageTimeResponse) SetHeaders(v map[string]*string) *ModifyUserSlsLogStorageTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserSlsLogStorageTimeResponse) SetStatusCode(v int32) *ModifyUserSlsLogStorageTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserSlsLogStorageTimeResponse) SetBody(v *ModifyUserSlsLogStorageTimeResponseBody) *ModifyUserSlsLogStorageTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyUserSlsLogStorageTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
