// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLineageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLineageInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetLineageInfoResponseBody) *GetLineageInfoResponse
	GetBody() *GetLineageInfoResponseBody
}

type GetLineageInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLineageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLineageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLineageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLineageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLineageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLineageInfoResponse) GetBody() *GetLineageInfoResponseBody {
	return s.Body
}

func (s *GetLineageInfoResponse) SetHeaders(v map[string]*string) *GetLineageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLineageInfoResponse) SetStatusCode(v int32) *GetLineageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLineageInfoResponse) SetBody(v *GetLineageInfoResponseBody) *GetLineageInfoResponse {
	s.Body = v
	return s
}

func (s *GetLineageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
