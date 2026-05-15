// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueInformationResponseBody) *GetQueueInformationResponse
	GetBody() *GetQueueInformationResponseBody
}

type GetQueueInformationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueInformationResponse) GoString() string {
	return s.String()
}

func (s *GetQueueInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueInformationResponse) GetBody() *GetQueueInformationResponseBody {
	return s.Body
}

func (s *GetQueueInformationResponse) SetHeaders(v map[string]*string) *GetQueueInformationResponse {
	s.Headers = v
	return s
}

func (s *GetQueueInformationResponse) SetStatusCode(v int32) *GetQueueInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueInformationResponse) SetBody(v *GetQueueInformationResponseBody) *GetQueueInformationResponse {
	s.Body = v
	return s
}

func (s *GetQueueInformationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
