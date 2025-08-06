// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveEditingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitLiveEditingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitLiveEditingResponse
	GetStatusCode() *int32
	SetBody(v *SubmitLiveEditingResponseBody) *SubmitLiveEditingResponse
	GetBody() *SubmitLiveEditingResponseBody
}

type SubmitLiveEditingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitLiveEditingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitLiveEditingResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveEditingResponse) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitLiveEditingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitLiveEditingResponse) GetBody() *SubmitLiveEditingResponseBody {
	return s.Body
}

func (s *SubmitLiveEditingResponse) SetHeaders(v map[string]*string) *SubmitLiveEditingResponse {
	s.Headers = v
	return s
}

func (s *SubmitLiveEditingResponse) SetStatusCode(v int32) *SubmitLiveEditingResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLiveEditingResponse) SetBody(v *SubmitLiveEditingResponseBody) *SubmitLiveEditingResponse {
	s.Body = v
	return s
}

func (s *SubmitLiveEditingResponse) Validate() error {
	return dara.Validate(s)
}
