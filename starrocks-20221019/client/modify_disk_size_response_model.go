// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskSizeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskSizeResponseBody) *ModifyDiskSizeResponse
	GetBody() *ModifyDiskSizeResponseBody
}

type ModifyDiskSizeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskSizeResponse) GetBody() *ModifyDiskSizeResponseBody {
	return s.Body
}

func (s *ModifyDiskSizeResponse) SetHeaders(v map[string]*string) *ModifyDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskSizeResponse) SetStatusCode(v int32) *ModifyDiskSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskSizeResponse) SetBody(v *ModifyDiskSizeResponseBody) *ModifyDiskSizeResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskSizeResponse) Validate() error {
	return dara.Validate(s)
}
