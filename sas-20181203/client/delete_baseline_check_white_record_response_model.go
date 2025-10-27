// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBaselineCheckWhiteRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBaselineCheckWhiteRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBaselineCheckWhiteRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBaselineCheckWhiteRecordResponseBody) *DeleteBaselineCheckWhiteRecordResponse
	GetBody() *DeleteBaselineCheckWhiteRecordResponseBody
}

type DeleteBaselineCheckWhiteRecordResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBaselineCheckWhiteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBaselineCheckWhiteRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBaselineCheckWhiteRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteBaselineCheckWhiteRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBaselineCheckWhiteRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBaselineCheckWhiteRecordResponse) GetBody() *DeleteBaselineCheckWhiteRecordResponseBody {
	return s.Body
}

func (s *DeleteBaselineCheckWhiteRecordResponse) SetHeaders(v map[string]*string) *DeleteBaselineCheckWhiteRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordResponse) SetStatusCode(v int32) *DeleteBaselineCheckWhiteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordResponse) SetBody(v *DeleteBaselineCheckWhiteRecordResponseBody) *DeleteBaselineCheckWhiteRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteBaselineCheckWhiteRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
