// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApInfoFromPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApInfoFromPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApInfoFromPoolResponse
	GetStatusCode() *int32
	SetBody(v *GetApInfoFromPoolResponseBody) *GetApInfoFromPoolResponse
	GetBody() *GetApInfoFromPoolResponseBody
}

type GetApInfoFromPoolResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApInfoFromPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApInfoFromPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApInfoFromPoolResponse) GoString() string {
	return s.String()
}

func (s *GetApInfoFromPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApInfoFromPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApInfoFromPoolResponse) GetBody() *GetApInfoFromPoolResponseBody {
	return s.Body
}

func (s *GetApInfoFromPoolResponse) SetHeaders(v map[string]*string) *GetApInfoFromPoolResponse {
	s.Headers = v
	return s
}

func (s *GetApInfoFromPoolResponse) SetStatusCode(v int32) *GetApInfoFromPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApInfoFromPoolResponse) SetBody(v *GetApInfoFromPoolResponseBody) *GetApInfoFromPoolResponse {
	s.Body = v
	return s
}

func (s *GetApInfoFromPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
