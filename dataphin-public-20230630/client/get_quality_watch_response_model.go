// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityWatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityWatchResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityWatchResponseBody) *GetQualityWatchResponse
	GetBody() *GetQualityWatchResponseBody
}

type GetQualityWatchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityWatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityWatchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchResponse) GoString() string {
	return s.String()
}

func (s *GetQualityWatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityWatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityWatchResponse) GetBody() *GetQualityWatchResponseBody {
	return s.Body
}

func (s *GetQualityWatchResponse) SetHeaders(v map[string]*string) *GetQualityWatchResponse {
	s.Headers = v
	return s
}

func (s *GetQualityWatchResponse) SetStatusCode(v int32) *GetQualityWatchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityWatchResponse) SetBody(v *GetQualityWatchResponseBody) *GetQualityWatchResponse {
	s.Body = v
	return s
}

func (s *GetQualityWatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
