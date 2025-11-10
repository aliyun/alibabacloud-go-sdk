// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersForParentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFoldersForParentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFoldersForParentResponse
	GetStatusCode() *int32
	SetBody(v *ListFoldersForParentResponseBody) *ListFoldersForParentResponse
	GetBody() *ListFoldersForParentResponseBody
}

type ListFoldersForParentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFoldersForParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFoldersForParentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentResponse) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFoldersForParentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFoldersForParentResponse) GetBody() *ListFoldersForParentResponseBody {
	return s.Body
}

func (s *ListFoldersForParentResponse) SetHeaders(v map[string]*string) *ListFoldersForParentResponse {
	s.Headers = v
	return s
}

func (s *ListFoldersForParentResponse) SetStatusCode(v int32) *ListFoldersForParentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFoldersForParentResponse) SetBody(v *ListFoldersForParentResponseBody) *ListFoldersForParentResponse {
	s.Body = v
	return s
}

func (s *ListFoldersForParentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
