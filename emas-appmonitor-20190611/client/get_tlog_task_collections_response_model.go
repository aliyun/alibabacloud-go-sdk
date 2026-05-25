// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogTaskCollectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTlogTaskCollectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTlogTaskCollectionsResponse
	GetStatusCode() *int32
	SetBody(v *GetTlogTaskCollectionsResponseBody) *GetTlogTaskCollectionsResponse
	GetBody() *GetTlogTaskCollectionsResponseBody
}

type GetTlogTaskCollectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTlogTaskCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTlogTaskCollectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTlogTaskCollectionsResponse) GoString() string {
	return s.String()
}

func (s *GetTlogTaskCollectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTlogTaskCollectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTlogTaskCollectionsResponse) GetBody() *GetTlogTaskCollectionsResponseBody {
	return s.Body
}

func (s *GetTlogTaskCollectionsResponse) SetHeaders(v map[string]*string) *GetTlogTaskCollectionsResponse {
	s.Headers = v
	return s
}

func (s *GetTlogTaskCollectionsResponse) SetStatusCode(v int32) *GetTlogTaskCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTlogTaskCollectionsResponse) SetBody(v *GetTlogTaskCollectionsResponseBody) *GetTlogTaskCollectionsResponse {
	s.Body = v
	return s
}

func (s *GetTlogTaskCollectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
