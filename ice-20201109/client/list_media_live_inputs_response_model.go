// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveInputsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaLiveInputsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaLiveInputsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaLiveInputsResponseBody) *ListMediaLiveInputsResponse
	GetBody() *ListMediaLiveInputsResponseBody
}

type ListMediaLiveInputsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaLiveInputsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaLiveInputsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaLiveInputsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaLiveInputsResponse) GetBody() *ListMediaLiveInputsResponseBody {
	return s.Body
}

func (s *ListMediaLiveInputsResponse) SetHeaders(v map[string]*string) *ListMediaLiveInputsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaLiveInputsResponse) SetStatusCode(v int32) *ListMediaLiveInputsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaLiveInputsResponse) SetBody(v *ListMediaLiveInputsResponseBody) *ListMediaLiveInputsResponse {
	s.Body = v
	return s
}

func (s *ListMediaLiveInputsResponse) Validate() error {
	return dara.Validate(s)
}
