// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessChannelOfStagingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessChannelOfStagingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessChannelOfStagingResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessChannelOfStagingResponseBody) *GetAccessChannelOfStagingResponse
	GetBody() *GetAccessChannelOfStagingResponseBody
}

type GetAccessChannelOfStagingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessChannelOfStagingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessChannelOfStagingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessChannelOfStagingResponse) GoString() string {
	return s.String()
}

func (s *GetAccessChannelOfStagingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessChannelOfStagingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessChannelOfStagingResponse) GetBody() *GetAccessChannelOfStagingResponseBody {
	return s.Body
}

func (s *GetAccessChannelOfStagingResponse) SetHeaders(v map[string]*string) *GetAccessChannelOfStagingResponse {
	s.Headers = v
	return s
}

func (s *GetAccessChannelOfStagingResponse) SetStatusCode(v int32) *GetAccessChannelOfStagingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessChannelOfStagingResponse) SetBody(v *GetAccessChannelOfStagingResponseBody) *GetAccessChannelOfStagingResponse {
	s.Body = v
	return s
}

func (s *GetAccessChannelOfStagingResponse) Validate() error {
	return dara.Validate(s)
}
