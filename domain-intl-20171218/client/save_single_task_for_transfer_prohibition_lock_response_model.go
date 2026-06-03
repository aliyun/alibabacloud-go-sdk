// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForTransferProhibitionLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForTransferProhibitionLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForTransferProhibitionLockResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForTransferProhibitionLockResponseBody) *SaveSingleTaskForTransferProhibitionLockResponse
	GetBody() *SaveSingleTaskForTransferProhibitionLockResponseBody
}

type SaveSingleTaskForTransferProhibitionLockResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForTransferProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForTransferProhibitionLockResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForTransferProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) GetBody() *SaveSingleTaskForTransferProhibitionLockResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForTransferProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) SetStatusCode(v int32) *SaveSingleTaskForTransferProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) SetBody(v *SaveSingleTaskForTransferProhibitionLockResponseBody) *SaveSingleTaskForTransferProhibitionLockResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
