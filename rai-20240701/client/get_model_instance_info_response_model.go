// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInstanceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelInstanceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelInstanceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetModelInstanceInfoResponseBody) *GetModelInstanceInfoResponse
	GetBody() *GetModelInstanceInfoResponseBody
}

type GetModelInstanceInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelInstanceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetModelInstanceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelInstanceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelInstanceInfoResponse) GetBody() *GetModelInstanceInfoResponseBody {
	return s.Body
}

func (s *GetModelInstanceInfoResponse) SetHeaders(v map[string]*string) *GetModelInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetModelInstanceInfoResponse) SetStatusCode(v int32) *GetModelInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelInstanceInfoResponse) SetBody(v *GetModelInstanceInfoResponseBody) *GetModelInstanceInfoResponse {
	s.Body = v
	return s
}

func (s *GetModelInstanceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
