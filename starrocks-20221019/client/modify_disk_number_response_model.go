// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskNumberResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskNumberResponseBody) *ModifyDiskNumberResponse
	GetBody() *ModifyDiskNumberResponseBody
}

type ModifyDiskNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskNumberResponse) GetBody() *ModifyDiskNumberResponseBody {
	return s.Body
}

func (s *ModifyDiskNumberResponse) SetHeaders(v map[string]*string) *ModifyDiskNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskNumberResponse) SetStatusCode(v int32) *ModifyDiskNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskNumberResponse) SetBody(v *ModifyDiskNumberResponseBody) *ModifyDiskNumberResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskNumberResponse) Validate() error {
	return dara.Validate(s)
}
