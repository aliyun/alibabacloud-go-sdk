// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemovalByFormInstanceIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRemovalByFormInstanceIdListResponse
	GetStatusCode() *int32
	SetBody(v *BatchRemovalByFormInstanceIdListResponseBody) *BatchRemovalByFormInstanceIdListResponse
	GetBody() *BatchRemovalByFormInstanceIdListResponseBody
}

type BatchRemovalByFormInstanceIdListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRemovalByFormInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRemovalByFormInstanceIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRemovalByFormInstanceIdListResponse) GetBody() *BatchRemovalByFormInstanceIdListResponseBody {
	return s.Body
}

func (s *BatchRemovalByFormInstanceIdListResponse) SetHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListResponse) SetStatusCode(v int32) *BatchRemovalByFormInstanceIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListResponse) SetBody(v *BatchRemovalByFormInstanceIdListResponseBody) *BatchRemovalByFormInstanceIdListResponse {
	s.Body = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
