// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPublishStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppPublishStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppPublishStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAppPublishStatusResponseBody) *GetAppPublishStatusResponse
	GetBody() *GetAppPublishStatusResponseBody
}

type GetAppPublishStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppPublishStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppPublishStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppPublishStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAppPublishStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppPublishStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppPublishStatusResponse) GetBody() *GetAppPublishStatusResponseBody {
	return s.Body
}

func (s *GetAppPublishStatusResponse) SetHeaders(v map[string]*string) *GetAppPublishStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAppPublishStatusResponse) SetStatusCode(v int32) *GetAppPublishStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppPublishStatusResponse) SetBody(v *GetAppPublishStatusResponseBody) *GetAppPublishStatusResponse {
	s.Body = v
	return s
}

func (s *GetAppPublishStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
