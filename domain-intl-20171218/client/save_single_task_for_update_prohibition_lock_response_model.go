// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForUpdateProhibitionLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForUpdateProhibitionLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForUpdateProhibitionLockResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForUpdateProhibitionLockResponseBody) *SaveSingleTaskForUpdateProhibitionLockResponse
	GetBody() *SaveSingleTaskForUpdateProhibitionLockResponseBody
}

type SaveSingleTaskForUpdateProhibitionLockResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForUpdateProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForUpdateProhibitionLockResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForUpdateProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) GetBody() *SaveSingleTaskForUpdateProhibitionLockResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForUpdateProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) SetStatusCode(v int32) *SaveSingleTaskForUpdateProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) SetBody(v *SaveSingleTaskForUpdateProhibitionLockResponseBody) *SaveSingleTaskForUpdateProhibitionLockResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
