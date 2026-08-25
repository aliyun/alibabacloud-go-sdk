// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageTestResultResponse
	GetStatusCode() *int32
	SetBody(v *GetImageTestResultResponseBody) *GetImageTestResultResponse
	GetBody() *GetImageTestResultResponseBody
}

type GetImageTestResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageTestResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageTestResultResponse) GetBody() *GetImageTestResultResponseBody {
	return s.Body
}

func (s *GetImageTestResultResponse) SetHeaders(v map[string]*string) *GetImageTestResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageTestResultResponse) SetStatusCode(v int32) *GetImageTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageTestResultResponse) SetBody(v *GetImageTestResultResponseBody) *GetImageTestResultResponse {
	s.Body = v
	return s
}

func (s *GetImageTestResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
