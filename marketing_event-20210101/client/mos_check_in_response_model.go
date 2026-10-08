// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMosCheckInResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MosCheckInResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MosCheckInResponse
	GetStatusCode() *int32
	SetBody(v *MosCheckInResponseBody) *MosCheckInResponse
	GetBody() *MosCheckInResponseBody
}

type MosCheckInResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MosCheckInResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MosCheckInResponse) String() string {
	return dara.Prettify(s)
}

func (s MosCheckInResponse) GoString() string {
	return s.String()
}

func (s *MosCheckInResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MosCheckInResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MosCheckInResponse) GetBody() *MosCheckInResponseBody {
	return s.Body
}

func (s *MosCheckInResponse) SetHeaders(v map[string]*string) *MosCheckInResponse {
	s.Headers = v
	return s
}

func (s *MosCheckInResponse) SetStatusCode(v int32) *MosCheckInResponse {
	s.StatusCode = &v
	return s
}

func (s *MosCheckInResponse) SetBody(v *MosCheckInResponseBody) *MosCheckInResponse {
	s.Body = v
	return s
}

func (s *MosCheckInResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
