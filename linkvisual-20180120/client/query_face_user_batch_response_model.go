// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceUserBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceUserBatchResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceUserBatchResponseBody) *QueryFaceUserBatchResponse
	GetBody() *QueryFaceUserBatchResponseBody
}

type QueryFaceUserBatchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserBatchResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceUserBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceUserBatchResponse) GetBody() *QueryFaceUserBatchResponseBody {
	return s.Body
}

func (s *QueryFaceUserBatchResponse) SetHeaders(v map[string]*string) *QueryFaceUserBatchResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserBatchResponse) SetStatusCode(v int32) *QueryFaceUserBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserBatchResponse) SetBody(v *QueryFaceUserBatchResponseBody) *QueryFaceUserBatchResponse {
	s.Body = v
	return s
}

func (s *QueryFaceUserBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
