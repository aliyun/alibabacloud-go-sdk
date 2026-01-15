// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineModelResponse
	GetStatusCode() *int32
	SetBody(v *OfflineModelResponseBody) *OfflineModelResponse
	GetBody() *OfflineModelResponseBody
}

type OfflineModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineModelResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineModelResponse) GoString() string {
	return s.String()
}

func (s *OfflineModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineModelResponse) GetBody() *OfflineModelResponseBody {
	return s.Body
}

func (s *OfflineModelResponse) SetHeaders(v map[string]*string) *OfflineModelResponse {
	s.Headers = v
	return s
}

func (s *OfflineModelResponse) SetStatusCode(v int32) *OfflineModelResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineModelResponse) SetBody(v *OfflineModelResponseBody) *OfflineModelResponse {
	s.Body = v
	return s
}

func (s *OfflineModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
