// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSitePauseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSitePauseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSitePauseResponse
	GetStatusCode() *int32
	SetBody(v *GetSitePauseResponseBody) *GetSitePauseResponse
	GetBody() *GetSitePauseResponseBody
}

type GetSitePauseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSitePauseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSitePauseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSitePauseResponse) GoString() string {
	return s.String()
}

func (s *GetSitePauseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSitePauseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSitePauseResponse) GetBody() *GetSitePauseResponseBody {
	return s.Body
}

func (s *GetSitePauseResponse) SetHeaders(v map[string]*string) *GetSitePauseResponse {
	s.Headers = v
	return s
}

func (s *GetSitePauseResponse) SetStatusCode(v int32) *GetSitePauseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSitePauseResponse) SetBody(v *GetSitePauseResponseBody) *GetSitePauseResponse {
	s.Body = v
	return s
}

func (s *GetSitePauseResponse) Validate() error {
	return dara.Validate(s)
}
