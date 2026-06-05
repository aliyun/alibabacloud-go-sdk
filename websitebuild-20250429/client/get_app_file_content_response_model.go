// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppFileContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppFileContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppFileContentResponse
	GetStatusCode() *int32
	SetBody(v *GetAppFileContentResponseBody) *GetAppFileContentResponse
	GetBody() *GetAppFileContentResponseBody
}

type GetAppFileContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppFileContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppFileContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppFileContentResponse) GoString() string {
	return s.String()
}

func (s *GetAppFileContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppFileContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppFileContentResponse) GetBody() *GetAppFileContentResponseBody {
	return s.Body
}

func (s *GetAppFileContentResponse) SetHeaders(v map[string]*string) *GetAppFileContentResponse {
	s.Headers = v
	return s
}

func (s *GetAppFileContentResponse) SetStatusCode(v int32) *GetAppFileContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppFileContentResponse) SetBody(v *GetAppFileContentResponseBody) *GetAppFileContentResponse {
	s.Body = v
	return s
}

func (s *GetAppFileContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
