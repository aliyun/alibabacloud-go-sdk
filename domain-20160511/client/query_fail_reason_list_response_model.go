// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFailReasonListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFailReasonListResponse
	GetStatusCode() *int32
	SetBody(v *QueryFailReasonListResponseBody) *QueryFailReasonListResponse
	GetBody() *QueryFailReasonListResponseBody
}

type QueryFailReasonListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFailReasonListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFailReasonListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonListResponse) GoString() string {
	return s.String()
}

func (s *QueryFailReasonListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFailReasonListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFailReasonListResponse) GetBody() *QueryFailReasonListResponseBody {
	return s.Body
}

func (s *QueryFailReasonListResponse) SetHeaders(v map[string]*string) *QueryFailReasonListResponse {
	s.Headers = v
	return s
}

func (s *QueryFailReasonListResponse) SetStatusCode(v int32) *QueryFailReasonListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailReasonListResponse) SetBody(v *QueryFailReasonListResponseBody) *QueryFailReasonListResponse {
	s.Body = v
	return s
}

func (s *QueryFailReasonListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
