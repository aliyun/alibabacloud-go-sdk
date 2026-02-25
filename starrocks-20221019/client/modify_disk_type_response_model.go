// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskTypeResponseBody) *ModifyDiskTypeResponse
	GetBody() *ModifyDiskTypeResponseBody
}

type ModifyDiskTypeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskTypeResponse) GetBody() *ModifyDiskTypeResponseBody {
	return s.Body
}

func (s *ModifyDiskTypeResponse) SetHeaders(v map[string]*string) *ModifyDiskTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskTypeResponse) SetStatusCode(v int32) *ModifyDiskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskTypeResponse) SetBody(v *ModifyDiskTypeResponseBody) *ModifyDiskTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
