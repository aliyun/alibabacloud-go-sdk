// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordStorageLifeCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutRecordStorageLifeCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutRecordStorageLifeCycleResponse
	GetStatusCode() *int32
	SetBody(v *PutRecordStorageLifeCycleResponseBody) *PutRecordStorageLifeCycleResponse
	GetBody() *PutRecordStorageLifeCycleResponseBody
}

type PutRecordStorageLifeCycleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutRecordStorageLifeCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutRecordStorageLifeCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutRecordStorageLifeCycleResponse) GoString() string {
	return s.String()
}

func (s *PutRecordStorageLifeCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutRecordStorageLifeCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutRecordStorageLifeCycleResponse) GetBody() *PutRecordStorageLifeCycleResponseBody {
	return s.Body
}

func (s *PutRecordStorageLifeCycleResponse) SetHeaders(v map[string]*string) *PutRecordStorageLifeCycleResponse {
	s.Headers = v
	return s
}

func (s *PutRecordStorageLifeCycleResponse) SetStatusCode(v int32) *PutRecordStorageLifeCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutRecordStorageLifeCycleResponse) SetBody(v *PutRecordStorageLifeCycleResponseBody) *PutRecordStorageLifeCycleResponse {
	s.Body = v
	return s
}

func (s *PutRecordStorageLifeCycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
