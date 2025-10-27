// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClearLogstoreStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClearLogstoreStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClearLogstoreStorageResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClearLogstoreStorageResponseBody) *ModifyClearLogstoreStorageResponse
	GetBody() *ModifyClearLogstoreStorageResponseBody
}

type ModifyClearLogstoreStorageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClearLogstoreStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClearLogstoreStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClearLogstoreStorageResponse) GoString() string {
	return s.String()
}

func (s *ModifyClearLogstoreStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClearLogstoreStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClearLogstoreStorageResponse) GetBody() *ModifyClearLogstoreStorageResponseBody {
	return s.Body
}

func (s *ModifyClearLogstoreStorageResponse) SetHeaders(v map[string]*string) *ModifyClearLogstoreStorageResponse {
	s.Headers = v
	return s
}

func (s *ModifyClearLogstoreStorageResponse) SetStatusCode(v int32) *ModifyClearLogstoreStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClearLogstoreStorageResponse) SetBody(v *ModifyClearLogstoreStorageResponseBody) *ModifyClearLogstoreStorageResponse {
	s.Body = v
	return s
}

func (s *ModifyClearLogstoreStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
