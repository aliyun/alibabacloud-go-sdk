// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorksEmbedListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorksEmbedListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorksEmbedListResponse
	GetStatusCode() *int32
	SetBody(v *GetWorksEmbedListResponseBody) *GetWorksEmbedListResponse
	GetBody() *GetWorksEmbedListResponseBody
}

type GetWorksEmbedListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorksEmbedListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorksEmbedListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorksEmbedListResponse) GoString() string {
	return s.String()
}

func (s *GetWorksEmbedListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorksEmbedListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorksEmbedListResponse) GetBody() *GetWorksEmbedListResponseBody {
	return s.Body
}

func (s *GetWorksEmbedListResponse) SetHeaders(v map[string]*string) *GetWorksEmbedListResponse {
	s.Headers = v
	return s
}

func (s *GetWorksEmbedListResponse) SetStatusCode(v int32) *GetWorksEmbedListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorksEmbedListResponse) SetBody(v *GetWorksEmbedListResponseBody) *GetWorksEmbedListResponse {
	s.Body = v
	return s
}

func (s *GetWorksEmbedListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
