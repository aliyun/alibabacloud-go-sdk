// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupLiveInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGroupLiveInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGroupLiveInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryGroupLiveInfoResponseBody) *QueryGroupLiveInfoResponse
	GetBody() *QueryGroupLiveInfoResponseBody
}

type QueryGroupLiveInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupLiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupLiveInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGroupLiveInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGroupLiveInfoResponse) GetBody() *QueryGroupLiveInfoResponseBody {
	return s.Body
}

func (s *QueryGroupLiveInfoResponse) SetHeaders(v map[string]*string) *QueryGroupLiveInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupLiveInfoResponse) SetStatusCode(v int32) *QueryGroupLiveInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupLiveInfoResponse) SetBody(v *QueryGroupLiveInfoResponseBody) *QueryGroupLiveInfoResponse {
	s.Body = v
	return s
}

func (s *QueryGroupLiveInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
