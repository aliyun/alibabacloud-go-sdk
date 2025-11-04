// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveInputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaLiveInputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaLiveInputResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaLiveInputResponseBody) *DeleteMediaLiveInputResponse
	GetBody() *DeleteMediaLiveInputResponseBody
}

type DeleteMediaLiveInputResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaLiveInputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaLiveInputResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveInputResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveInputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaLiveInputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaLiveInputResponse) GetBody() *DeleteMediaLiveInputResponseBody {
	return s.Body
}

func (s *DeleteMediaLiveInputResponse) SetHeaders(v map[string]*string) *DeleteMediaLiveInputResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaLiveInputResponse) SetStatusCode(v int32) *DeleteMediaLiveInputResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaLiveInputResponse) SetBody(v *DeleteMediaLiveInputResponseBody) *DeleteMediaLiveInputResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaLiveInputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
