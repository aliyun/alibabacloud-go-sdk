// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderActivateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderActivateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderActivateResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCreatingOrderActivateResponseBody) *SaveSingleTaskForCreatingOrderActivateResponse
	GetBody() *SaveSingleTaskForCreatingOrderActivateResponseBody
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCreatingOrderActivateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderActivateResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderActivateResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) GetBody() *SaveSingleTaskForCreatingOrderActivateResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderActivateResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderActivateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) SetBody(v *SaveSingleTaskForCreatingOrderActivateResponseBody) *SaveSingleTaskForCreatingOrderActivateResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
