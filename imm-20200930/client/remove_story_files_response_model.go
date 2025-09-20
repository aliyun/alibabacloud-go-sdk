// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveStoryFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveStoryFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveStoryFilesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveStoryFilesResponseBody) *RemoveStoryFilesResponse
	GetBody() *RemoveStoryFilesResponseBody
}

type RemoveStoryFilesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveStoryFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveStoryFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesResponse) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveStoryFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveStoryFilesResponse) GetBody() *RemoveStoryFilesResponseBody {
	return s.Body
}

func (s *RemoveStoryFilesResponse) SetHeaders(v map[string]*string) *RemoveStoryFilesResponse {
	s.Headers = v
	return s
}

func (s *RemoveStoryFilesResponse) SetStatusCode(v int32) *RemoveStoryFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveStoryFilesResponse) SetBody(v *RemoveStoryFilesResponseBody) *RemoveStoryFilesResponse {
	s.Body = v
	return s
}

func (s *RemoveStoryFilesResponse) Validate() error {
	return dara.Validate(s)
}
