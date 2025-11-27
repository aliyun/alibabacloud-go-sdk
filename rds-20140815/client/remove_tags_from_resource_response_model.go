// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsFromResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTagsFromResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTagsFromResourceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTagsFromResourceResponseBody) *RemoveTagsFromResourceResponse
	GetBody() *RemoveTagsFromResourceResponseBody
}

type RemoveTagsFromResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTagsFromResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTagsFromResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsFromResourceResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsFromResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTagsFromResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTagsFromResourceResponse) GetBody() *RemoveTagsFromResourceResponseBody {
	return s.Body
}

func (s *RemoveTagsFromResourceResponse) SetHeaders(v map[string]*string) *RemoveTagsFromResourceResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsFromResourceResponse) SetStatusCode(v int32) *RemoveTagsFromResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagsFromResourceResponse) SetBody(v *RemoveTagsFromResourceResponseBody) *RemoveTagsFromResourceResponse {
	s.Body = v
	return s
}

func (s *RemoveTagsFromResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
