// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMosQueryRegistrationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MosQueryRegistrationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MosQueryRegistrationInfoResponse
	GetStatusCode() *int32
	SetBody(v *MosQueryRegistrationInfoResponseBody) *MosQueryRegistrationInfoResponse
	GetBody() *MosQueryRegistrationInfoResponseBody
}

type MosQueryRegistrationInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MosQueryRegistrationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MosQueryRegistrationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s MosQueryRegistrationInfoResponse) GoString() string {
	return s.String()
}

func (s *MosQueryRegistrationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MosQueryRegistrationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MosQueryRegistrationInfoResponse) GetBody() *MosQueryRegistrationInfoResponseBody {
	return s.Body
}

func (s *MosQueryRegistrationInfoResponse) SetHeaders(v map[string]*string) *MosQueryRegistrationInfoResponse {
	s.Headers = v
	return s
}

func (s *MosQueryRegistrationInfoResponse) SetStatusCode(v int32) *MosQueryRegistrationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *MosQueryRegistrationInfoResponse) SetBody(v *MosQueryRegistrationInfoResponseBody) *MosQueryRegistrationInfoResponse {
	s.Body = v
	return s
}

func (s *MosQueryRegistrationInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
