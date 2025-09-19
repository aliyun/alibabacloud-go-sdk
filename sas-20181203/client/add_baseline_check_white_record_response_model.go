// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaselineCheckWhiteRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBaselineCheckWhiteRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBaselineCheckWhiteRecordResponse
	GetStatusCode() *int32
	SetBody(v *AddBaselineCheckWhiteRecordResponseBody) *AddBaselineCheckWhiteRecordResponse
	GetBody() *AddBaselineCheckWhiteRecordResponseBody
}

type AddBaselineCheckWhiteRecordResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBaselineCheckWhiteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBaselineCheckWhiteRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBaselineCheckWhiteRecordResponse) GoString() string {
	return s.String()
}

func (s *AddBaselineCheckWhiteRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBaselineCheckWhiteRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBaselineCheckWhiteRecordResponse) GetBody() *AddBaselineCheckWhiteRecordResponseBody {
	return s.Body
}

func (s *AddBaselineCheckWhiteRecordResponse) SetHeaders(v map[string]*string) *AddBaselineCheckWhiteRecordResponse {
	s.Headers = v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponse) SetStatusCode(v int32) *AddBaselineCheckWhiteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponse) SetBody(v *AddBaselineCheckWhiteRecordResponseBody) *AddBaselineCheckWhiteRecordResponse {
	s.Body = v
	return s
}

func (s *AddBaselineCheckWhiteRecordResponse) Validate() error {
	return dara.Validate(s)
}
