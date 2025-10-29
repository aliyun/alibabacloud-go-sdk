// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpApiOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpApiOperationResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpApiOperationResponseBody) *GetHttpApiOperationResponse
	GetBody() *GetHttpApiOperationResponseBody
}

type GetHttpApiOperationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpApiOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpApiOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiOperationResponse) GoString() string {
	return s.String()
}

func (s *GetHttpApiOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpApiOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpApiOperationResponse) GetBody() *GetHttpApiOperationResponseBody {
	return s.Body
}

func (s *GetHttpApiOperationResponse) SetHeaders(v map[string]*string) *GetHttpApiOperationResponse {
	s.Headers = v
	return s
}

func (s *GetHttpApiOperationResponse) SetStatusCode(v int32) *GetHttpApiOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpApiOperationResponse) SetBody(v *GetHttpApiOperationResponseBody) *GetHttpApiOperationResponse {
	s.Body = v
	return s
}

func (s *GetHttpApiOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
