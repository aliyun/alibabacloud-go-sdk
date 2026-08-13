// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySyncResultResponse
	GetStatusCode() *int32
	SetBody(v *QuerySyncResultResponseBody) *QuerySyncResultResponse
	GetBody() *QuerySyncResultResponseBody
}

type QuerySyncResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncResultResponse) GoString() string {
	return s.String()
}

func (s *QuerySyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySyncResultResponse) GetBody() *QuerySyncResultResponseBody {
	return s.Body
}

func (s *QuerySyncResultResponse) SetHeaders(v map[string]*string) *QuerySyncResultResponse {
	s.Headers = v
	return s
}

func (s *QuerySyncResultResponse) SetStatusCode(v int32) *QuerySyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySyncResultResponse) SetBody(v *QuerySyncResultResponseBody) *QuerySyncResultResponse {
	s.Body = v
	return s
}

func (s *QuerySyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
