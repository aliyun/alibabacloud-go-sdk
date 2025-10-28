// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveImageLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveImageLabelsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveImageLabelsResponseBody) *RemoveImageLabelsResponse
	GetBody() *RemoveImageLabelsResponseBody
}

type RemoveImageLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveImageLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveImageLabelsResponse) GetBody() *RemoveImageLabelsResponseBody {
	return s.Body
}

func (s *RemoveImageLabelsResponse) SetHeaders(v map[string]*string) *RemoveImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageLabelsResponse) SetStatusCode(v int32) *RemoveImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageLabelsResponse) SetBody(v *RemoveImageLabelsResponseBody) *RemoveImageLabelsResponse {
	s.Body = v
	return s
}

func (s *RemoveImageLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
