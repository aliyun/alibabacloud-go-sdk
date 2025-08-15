// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrailStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrailStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrailStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTrailStatusResponseBody) *GetTrailStatusResponse
	GetBody() *GetTrailStatusResponseBody
}

type GetTrailStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrailStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrailStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrailStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTrailStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrailStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrailStatusResponse) GetBody() *GetTrailStatusResponseBody {
	return s.Body
}

func (s *GetTrailStatusResponse) SetHeaders(v map[string]*string) *GetTrailStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTrailStatusResponse) SetStatusCode(v int32) *GetTrailStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrailStatusResponse) SetBody(v *GetTrailStatusResponseBody) *GetTrailStatusResponse {
	s.Body = v
	return s
}

func (s *GetTrailStatusResponse) Validate() error {
	return dara.Validate(s)
}
