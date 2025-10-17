// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDesigateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDesigateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDesigateInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDesigateInfoResponseBody) *GetDesigateInfoResponse
	GetBody() *GetDesigateInfoResponseBody
}

type GetDesigateInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDesigateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDesigateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDesigateInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDesigateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDesigateInfoResponse) GetBody() *GetDesigateInfoResponseBody {
	return s.Body
}

func (s *GetDesigateInfoResponse) SetHeaders(v map[string]*string) *GetDesigateInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDesigateInfoResponse) SetStatusCode(v int32) *GetDesigateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDesigateInfoResponse) SetBody(v *GetDesigateInfoResponseBody) *GetDesigateInfoResponse {
	s.Body = v
	return s
}

func (s *GetDesigateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
