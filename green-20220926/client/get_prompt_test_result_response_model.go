// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPromptTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPromptTestResultResponse
	GetStatusCode() *int32
	SetBody(v *GetPromptTestResultResponseBody) *GetPromptTestResultResponse
	GetBody() *GetPromptTestResultResponseBody
}

type GetPromptTestResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPromptTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPromptTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTestResultResponse) GoString() string {
	return s.String()
}

func (s *GetPromptTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPromptTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPromptTestResultResponse) GetBody() *GetPromptTestResultResponseBody {
	return s.Body
}

func (s *GetPromptTestResultResponse) SetHeaders(v map[string]*string) *GetPromptTestResultResponse {
	s.Headers = v
	return s
}

func (s *GetPromptTestResultResponse) SetStatusCode(v int32) *GetPromptTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPromptTestResultResponse) SetBody(v *GetPromptTestResultResponseBody) *GetPromptTestResultResponse {
	s.Body = v
	return s
}

func (s *GetPromptTestResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
