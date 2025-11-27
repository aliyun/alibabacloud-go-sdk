// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsToResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTagsToResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTagsToResourceResponse
	GetStatusCode() *int32
	SetBody(v *AddTagsToResourceResponseBody) *AddTagsToResourceResponse
	GetBody() *AddTagsToResourceResponseBody
}

type AddTagsToResourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTagsToResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTagsToResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTagsToResourceResponse) GoString() string {
	return s.String()
}

func (s *AddTagsToResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTagsToResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTagsToResourceResponse) GetBody() *AddTagsToResourceResponseBody {
	return s.Body
}

func (s *AddTagsToResourceResponse) SetHeaders(v map[string]*string) *AddTagsToResourceResponse {
	s.Headers = v
	return s
}

func (s *AddTagsToResourceResponse) SetStatusCode(v int32) *AddTagsToResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsToResourceResponse) SetBody(v *AddTagsToResourceResponseBody) *AddTagsToResourceResponse {
	s.Body = v
	return s
}

func (s *AddTagsToResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
