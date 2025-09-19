// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomBlockRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCustomBlockRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCustomBlockRecordResponse
	GetStatusCode() *int32
	SetBody(v *DisableCustomBlockRecordResponseBody) *DisableCustomBlockRecordResponse
	GetBody() *DisableCustomBlockRecordResponseBody
}

type DisableCustomBlockRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCustomBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCustomBlockRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *DisableCustomBlockRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCustomBlockRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCustomBlockRecordResponse) GetBody() *DisableCustomBlockRecordResponseBody {
	return s.Body
}

func (s *DisableCustomBlockRecordResponse) SetHeaders(v map[string]*string) *DisableCustomBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *DisableCustomBlockRecordResponse) SetStatusCode(v int32) *DisableCustomBlockRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCustomBlockRecordResponse) SetBody(v *DisableCustomBlockRecordResponseBody) *DisableCustomBlockRecordResponse {
	s.Body = v
	return s
}

func (s *DisableCustomBlockRecordResponse) Validate() error {
	return dara.Validate(s)
}
