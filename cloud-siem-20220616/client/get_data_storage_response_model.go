// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataStorageResponse
	GetStatusCode() *int32
	SetBody(v *GetDataStorageResponseBody) *GetDataStorageResponse
	GetBody() *GetDataStorageResponseBody
}

type GetDataStorageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageResponse) GoString() string {
	return s.String()
}

func (s *GetDataStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataStorageResponse) GetBody() *GetDataStorageResponseBody {
	return s.Body
}

func (s *GetDataStorageResponse) SetHeaders(v map[string]*string) *GetDataStorageResponse {
	s.Headers = v
	return s
}

func (s *GetDataStorageResponse) SetStatusCode(v int32) *GetDataStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataStorageResponse) SetBody(v *GetDataStorageResponseBody) *GetDataStorageResponse {
	s.Body = v
	return s
}

func (s *GetDataStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
