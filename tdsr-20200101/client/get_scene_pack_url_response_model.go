// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePackUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScenePackUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScenePackUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetScenePackUrlResponseBody) *GetScenePackUrlResponse
	GetBody() *GetScenePackUrlResponseBody
}

type GetScenePackUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScenePackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScenePackUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScenePackUrlResponse) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScenePackUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScenePackUrlResponse) GetBody() *GetScenePackUrlResponseBody {
	return s.Body
}

func (s *GetScenePackUrlResponse) SetHeaders(v map[string]*string) *GetScenePackUrlResponse {
	s.Headers = v
	return s
}

func (s *GetScenePackUrlResponse) SetStatusCode(v int32) *GetScenePackUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePackUrlResponse) SetBody(v *GetScenePackUrlResponseBody) *GetScenePackUrlResponse {
	s.Body = v
	return s
}

func (s *GetScenePackUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
