// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineCheckWhiteRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaselineCheckWhiteRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaselineCheckWhiteRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListBaselineCheckWhiteRecordResponseBody) *ListBaselineCheckWhiteRecordResponse
	GetBody() *ListBaselineCheckWhiteRecordResponseBody
}

type ListBaselineCheckWhiteRecordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaselineCheckWhiteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaselineCheckWhiteRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineCheckWhiteRecordResponse) GoString() string {
	return s.String()
}

func (s *ListBaselineCheckWhiteRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaselineCheckWhiteRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaselineCheckWhiteRecordResponse) GetBody() *ListBaselineCheckWhiteRecordResponseBody {
	return s.Body
}

func (s *ListBaselineCheckWhiteRecordResponse) SetHeaders(v map[string]*string) *ListBaselineCheckWhiteRecordResponse {
	s.Headers = v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponse) SetStatusCode(v int32) *ListBaselineCheckWhiteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponse) SetBody(v *ListBaselineCheckWhiteRecordResponseBody) *ListBaselineCheckWhiteRecordResponse {
	s.Body = v
	return s
}

func (s *ListBaselineCheckWhiteRecordResponse) Validate() error {
	return dara.Validate(s)
}
