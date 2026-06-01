// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayHistoryServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRayHistoryServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRayHistoryServerResponse
	GetStatusCode() *int32
	SetBody(v *GetRayHistoryServerResponseBody) *GetRayHistoryServerResponse
	GetBody() *GetRayHistoryServerResponseBody
}

type GetRayHistoryServerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRayHistoryServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRayHistoryServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRayHistoryServerResponse) GoString() string {
	return s.String()
}

func (s *GetRayHistoryServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRayHistoryServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRayHistoryServerResponse) GetBody() *GetRayHistoryServerResponseBody {
	return s.Body
}

func (s *GetRayHistoryServerResponse) SetHeaders(v map[string]*string) *GetRayHistoryServerResponse {
	s.Headers = v
	return s
}

func (s *GetRayHistoryServerResponse) SetStatusCode(v int32) *GetRayHistoryServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRayHistoryServerResponse) SetBody(v *GetRayHistoryServerResponseBody) *GetRayHistoryServerResponse {
	s.Body = v
	return s
}

func (s *GetRayHistoryServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
