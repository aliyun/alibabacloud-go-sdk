// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveInstanceWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveInstanceWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *RemoveInstanceWhiteListResponseBody) *RemoveInstanceWhiteListResponse
	GetBody() *RemoveInstanceWhiteListResponseBody
}

type RemoveInstanceWhiteListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstanceWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstanceWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveInstanceWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveInstanceWhiteListResponse) GetBody() *RemoveInstanceWhiteListResponseBody {
	return s.Body
}

func (s *RemoveInstanceWhiteListResponse) SetHeaders(v map[string]*string) *RemoveInstanceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstanceWhiteListResponse) SetStatusCode(v int32) *RemoveInstanceWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceWhiteListResponse) SetBody(v *RemoveInstanceWhiteListResponseBody) *RemoveInstanceWhiteListResponse {
	s.Body = v
	return s
}

func (s *RemoveInstanceWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
