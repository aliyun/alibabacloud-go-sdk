// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertWhiteListSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertWhiteListSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertWhiteListSettingResponse
	GetStatusCode() *int32
	SetBody(v *InsertWhiteListSettingResponseBody) *InsertWhiteListSettingResponse
	GetBody() *InsertWhiteListSettingResponseBody
}

type InsertWhiteListSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertWhiteListSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertWhiteListSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertWhiteListSettingResponse) GoString() string {
	return s.String()
}

func (s *InsertWhiteListSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertWhiteListSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertWhiteListSettingResponse) GetBody() *InsertWhiteListSettingResponseBody {
	return s.Body
}

func (s *InsertWhiteListSettingResponse) SetHeaders(v map[string]*string) *InsertWhiteListSettingResponse {
	s.Headers = v
	return s
}

func (s *InsertWhiteListSettingResponse) SetStatusCode(v int32) *InsertWhiteListSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertWhiteListSettingResponse) SetBody(v *InsertWhiteListSettingResponseBody) *InsertWhiteListSettingResponse {
	s.Body = v
	return s
}

func (s *InsertWhiteListSettingResponse) Validate() error {
	return dara.Validate(s)
}
