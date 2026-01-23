// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityWatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityWatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityWatchTaskResponseBody) *GetQualityWatchTaskResponse
	GetBody() *GetQualityWatchTaskResponseBody
}

type GetQualityWatchTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityWatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityWatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchTaskResponse) GoString() string {
	return s.String()
}

func (s *GetQualityWatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityWatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityWatchTaskResponse) GetBody() *GetQualityWatchTaskResponseBody {
	return s.Body
}

func (s *GetQualityWatchTaskResponse) SetHeaders(v map[string]*string) *GetQualityWatchTaskResponse {
	s.Headers = v
	return s
}

func (s *GetQualityWatchTaskResponse) SetStatusCode(v int32) *GetQualityWatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityWatchTaskResponse) SetBody(v *GetQualityWatchTaskResponseBody) *GetQualityWatchTaskResponse {
	s.Body = v
	return s
}

func (s *GetQualityWatchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
