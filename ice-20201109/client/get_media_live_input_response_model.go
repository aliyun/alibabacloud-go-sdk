// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaLiveInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaLiveInputResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaLiveInputResponseBody) *GetMediaLiveInputResponse
	GetBody() *GetMediaLiveInputResponseBody
}

type GetMediaLiveInputResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaLiveInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaLiveInputResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputResponse) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaLiveInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaLiveInputResponse) GetBody() *GetMediaLiveInputResponseBody {
	return s.Body
}

func (s *GetMediaLiveInputResponse) SetHeaders(v map[string]*string) *GetMediaLiveInputResponse {
	s.Headers = v
	return s
}

func (s *GetMediaLiveInputResponse) SetStatusCode(v int32) *GetMediaLiveInputResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaLiveInputResponse) SetBody(v *GetMediaLiveInputResponseBody) *GetMediaLiveInputResponse {
	s.Body = v
	return s
}

func (s *GetMediaLiveInputResponse) Validate() error {
	return dara.Validate(s)
}
