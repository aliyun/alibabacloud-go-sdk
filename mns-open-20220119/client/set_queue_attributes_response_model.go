// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetQueueAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetQueueAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetQueueAttributesResponse
	GetStatusCode() *int32
	SetBody(v *SetQueueAttributesResponseBody) *SetQueueAttributesResponse
	GetBody() *SetQueueAttributesResponseBody
}

type SetQueueAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetQueueAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetQueueAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetQueueAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetQueueAttributesResponse) GetBody() *SetQueueAttributesResponseBody {
	return s.Body
}

func (s *SetQueueAttributesResponse) SetHeaders(v map[string]*string) *SetQueueAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetQueueAttributesResponse) SetStatusCode(v int32) *SetQueueAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetQueueAttributesResponse) SetBody(v *SetQueueAttributesResponseBody) *SetQueueAttributesResponse {
	s.Body = v
	return s
}

func (s *SetQueueAttributesResponse) Validate() error {
	return dara.Validate(s)
}
