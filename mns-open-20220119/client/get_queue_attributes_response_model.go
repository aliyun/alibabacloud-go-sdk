// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueAttributesResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueAttributesResponseBody) *GetQueueAttributesResponse
	GetBody() *GetQueueAttributesResponseBody
}

type GetQueueAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueAttributesResponse) GetBody() *GetQueueAttributesResponseBody {
	return s.Body
}

func (s *GetQueueAttributesResponse) SetHeaders(v map[string]*string) *GetQueueAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetQueueAttributesResponse) SetStatusCode(v int32) *GetQueueAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueAttributesResponse) SetBody(v *GetQueueAttributesResponseBody) *GetQueueAttributesResponse {
	s.Body = v
	return s
}

func (s *GetQueueAttributesResponse) Validate() error {
	return dara.Validate(s)
}
