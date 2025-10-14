// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDataStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDataStorageResponse
	GetStatusCode() *int32
	SetBody(v *ResetDataStorageResponseBody) *ResetDataStorageResponse
	GetBody() *ResetDataStorageResponseBody
}

type ResetDataStorageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDataStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDataStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDataStorageResponse) GoString() string {
	return s.String()
}

func (s *ResetDataStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDataStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDataStorageResponse) GetBody() *ResetDataStorageResponseBody {
	return s.Body
}

func (s *ResetDataStorageResponse) SetHeaders(v map[string]*string) *ResetDataStorageResponse {
	s.Headers = v
	return s
}

func (s *ResetDataStorageResponse) SetStatusCode(v int32) *ResetDataStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDataStorageResponse) SetBody(v *ResetDataStorageResponseBody) *ResetDataStorageResponse {
	s.Body = v
	return s
}

func (s *ResetDataStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
