// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTestResultListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTestResultListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTestResultListResponse
	GetStatusCode() *int32
	SetBody(v *GetTestResultListResponseBody) *GetTestResultListResponse
	GetBody() *GetTestResultListResponseBody
}

type GetTestResultListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTestResultListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTestResultListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListResponse) GoString() string {
	return s.String()
}

func (s *GetTestResultListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTestResultListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTestResultListResponse) GetBody() *GetTestResultListResponseBody {
	return s.Body
}

func (s *GetTestResultListResponse) SetHeaders(v map[string]*string) *GetTestResultListResponse {
	s.Headers = v
	return s
}

func (s *GetTestResultListResponse) SetStatusCode(v int32) *GetTestResultListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTestResultListResponse) SetBody(v *GetTestResultListResponseBody) *GetTestResultListResponse {
	s.Body = v
	return s
}

func (s *GetTestResultListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
