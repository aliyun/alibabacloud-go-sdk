// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoBuildRecordStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoBuildRecordStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoBuildRecordStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoBuildRecordStatusResponseBody) *GetRepoBuildRecordStatusResponse
	GetBody() *GetRepoBuildRecordStatusResponseBody
}

type GetRepoBuildRecordStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoBuildRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoBuildRecordStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoBuildRecordStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoBuildRecordStatusResponse) GetBody() *GetRepoBuildRecordStatusResponseBody {
	return s.Body
}

func (s *GetRepoBuildRecordStatusResponse) SetHeaders(v map[string]*string) *GetRepoBuildRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRepoBuildRecordStatusResponse) SetStatusCode(v int32) *GetRepoBuildRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponse) SetBody(v *GetRepoBuildRecordStatusResponseBody) *GetRepoBuildRecordStatusResponse {
	s.Body = v
	return s
}

func (s *GetRepoBuildRecordStatusResponse) Validate() error {
	return dara.Validate(s)
}
