// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStaStatusListByApResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStaStatusListByApResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStaStatusListByApResponse
	GetStatusCode() *int32
	SetBody(v *GetStaStatusListByApResponseBody) *GetStaStatusListByApResponse
	GetBody() *GetStaStatusListByApResponseBody
}

type GetStaStatusListByApResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStaStatusListByApResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStaStatusListByApResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStaStatusListByApResponse) GoString() string {
	return s.String()
}

func (s *GetStaStatusListByApResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStaStatusListByApResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStaStatusListByApResponse) GetBody() *GetStaStatusListByApResponseBody {
	return s.Body
}

func (s *GetStaStatusListByApResponse) SetHeaders(v map[string]*string) *GetStaStatusListByApResponse {
	s.Headers = v
	return s
}

func (s *GetStaStatusListByApResponse) SetStatusCode(v int32) *GetStaStatusListByApResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStaStatusListByApResponse) SetBody(v *GetStaStatusListByApResponseBody) *GetStaStatusListByApResponse {
	s.Body = v
	return s
}

func (s *GetStaStatusListByApResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
