// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetC2paSignResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetC2paSignResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetC2paSignResultResponse
	GetStatusCode() *int32
	SetBody(v *GetC2paSignResultResponseBody) *GetC2paSignResultResponse
	GetBody() *GetC2paSignResultResponseBody
}

type GetC2paSignResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetC2paSignResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetC2paSignResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetC2paSignResultResponse) GoString() string {
	return s.String()
}

func (s *GetC2paSignResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetC2paSignResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetC2paSignResultResponse) GetBody() *GetC2paSignResultResponseBody {
	return s.Body
}

func (s *GetC2paSignResultResponse) SetHeaders(v map[string]*string) *GetC2paSignResultResponse {
	s.Headers = v
	return s
}

func (s *GetC2paSignResultResponse) SetStatusCode(v int32) *GetC2paSignResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetC2paSignResultResponse) SetBody(v *GetC2paSignResultResponseBody) *GetC2paSignResultResponse {
	s.Body = v
	return s
}

func (s *GetC2paSignResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
