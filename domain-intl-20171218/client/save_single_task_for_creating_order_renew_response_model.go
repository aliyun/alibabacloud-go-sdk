// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderRenewResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCreatingOrderRenewResponseBody) *SaveSingleTaskForCreatingOrderRenewResponse
	GetBody() *SaveSingleTaskForCreatingOrderRenewResponseBody
}

type SaveSingleTaskForCreatingOrderRenewResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCreatingOrderRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRenewResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) GetBody() *SaveSingleTaskForCreatingOrderRenewResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderRenewResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) SetBody(v *SaveSingleTaskForCreatingOrderRenewResponseBody) *SaveSingleTaskForCreatingOrderRenewResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
