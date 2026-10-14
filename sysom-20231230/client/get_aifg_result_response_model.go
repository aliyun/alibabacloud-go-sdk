// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAifgResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAifgResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAifgResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAifgResultResponseBody) *GetAifgResultResponse
	GetBody() *GetAifgResultResponseBody
}

type GetAifgResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAifgResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAifgResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAifgResultResponse) GoString() string {
	return s.String()
}

func (s *GetAifgResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAifgResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAifgResultResponse) GetBody() *GetAifgResultResponseBody {
	return s.Body
}

func (s *GetAifgResultResponse) SetHeaders(v map[string]*string) *GetAifgResultResponse {
	s.Headers = v
	return s
}

func (s *GetAifgResultResponse) SetStatusCode(v int32) *GetAifgResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAifgResultResponse) SetBody(v *GetAifgResultResponseBody) *GetAifgResultResponse {
	s.Body = v
	return s
}

func (s *GetAifgResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
