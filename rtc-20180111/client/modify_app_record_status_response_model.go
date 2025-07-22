// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppRecordStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppRecordStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppRecordStatusResponseBody) *ModifyAppRecordStatusResponse
	GetBody() *ModifyAppRecordStatusResponseBody
}

type ModifyAppRecordStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppRecordStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppRecordStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppRecordStatusResponse) GetBody() *ModifyAppRecordStatusResponseBody {
	return s.Body
}

func (s *ModifyAppRecordStatusResponse) SetHeaders(v map[string]*string) *ModifyAppRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppRecordStatusResponse) SetStatusCode(v int32) *ModifyAppRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppRecordStatusResponse) SetBody(v *ModifyAppRecordStatusResponseBody) *ModifyAppRecordStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyAppRecordStatusResponse) Validate() error {
	return dara.Validate(s)
}
