// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConferenceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConferenceInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryConferenceInfoResponseBody) *QueryConferenceInfoResponse
	GetBody() *QueryConferenceInfoResponseBody
}

type QueryConferenceInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConferenceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConferenceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConferenceInfoResponse) GetBody() *QueryConferenceInfoResponseBody {
	return s.Body
}

func (s *QueryConferenceInfoResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoResponse) SetStatusCode(v int32) *QueryConferenceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceInfoResponse) SetBody(v *QueryConferenceInfoResponseBody) *QueryConferenceInfoResponse {
	s.Body = v
	return s
}

func (s *QueryConferenceInfoResponse) Validate() error {
	return dara.Validate(s)
}
