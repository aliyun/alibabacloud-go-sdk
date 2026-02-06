// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrServerInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsrServerInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsrServerInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAsrServerInfoResponseBody) *GetAsrServerInfoResponse
	GetBody() *GetAsrServerInfoResponseBody
}

type GetAsrServerInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsrServerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsrServerInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsrServerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAsrServerInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsrServerInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsrServerInfoResponse) GetBody() *GetAsrServerInfoResponseBody {
	return s.Body
}

func (s *GetAsrServerInfoResponse) SetHeaders(v map[string]*string) *GetAsrServerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAsrServerInfoResponse) SetStatusCode(v int32) *GetAsrServerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrServerInfoResponse) SetBody(v *GetAsrServerInfoResponseBody) *GetAsrServerInfoResponse {
	s.Body = v
	return s
}

func (s *GetAsrServerInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
