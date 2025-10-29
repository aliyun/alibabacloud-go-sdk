// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveShowFromShowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveShowFromShowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveShowFromShowListResponse
	GetStatusCode() *int32
	SetBody(v *RemoveShowFromShowListResponseBody) *RemoveShowFromShowListResponse
	GetBody() *RemoveShowFromShowListResponseBody
}

type RemoveShowFromShowListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveShowFromShowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveShowFromShowListResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveShowFromShowListResponse) GoString() string {
	return s.String()
}

func (s *RemoveShowFromShowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveShowFromShowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveShowFromShowListResponse) GetBody() *RemoveShowFromShowListResponseBody {
	return s.Body
}

func (s *RemoveShowFromShowListResponse) SetHeaders(v map[string]*string) *RemoveShowFromShowListResponse {
	s.Headers = v
	return s
}

func (s *RemoveShowFromShowListResponse) SetStatusCode(v int32) *RemoveShowFromShowListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveShowFromShowListResponse) SetBody(v *RemoveShowFromShowListResponseBody) *RemoveShowFromShowListResponse {
	s.Body = v
	return s
}

func (s *RemoveShowFromShowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
