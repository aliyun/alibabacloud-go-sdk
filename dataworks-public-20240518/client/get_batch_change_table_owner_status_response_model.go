// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchChangeTableOwnerStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchChangeTableOwnerStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchChangeTableOwnerStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchChangeTableOwnerStatusResponseBody) *GetBatchChangeTableOwnerStatusResponse
	GetBody() *GetBatchChangeTableOwnerStatusResponseBody
}

type GetBatchChangeTableOwnerStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchChangeTableOwnerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchChangeTableOwnerStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchChangeTableOwnerStatusResponse) GoString() string {
	return s.String()
}

func (s *GetBatchChangeTableOwnerStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchChangeTableOwnerStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchChangeTableOwnerStatusResponse) GetBody() *GetBatchChangeTableOwnerStatusResponseBody {
	return s.Body
}

func (s *GetBatchChangeTableOwnerStatusResponse) SetHeaders(v map[string]*string) *GetBatchChangeTableOwnerStatusResponse {
	s.Headers = v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponse) SetStatusCode(v int32) *GetBatchChangeTableOwnerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponse) SetBody(v *GetBatchChangeTableOwnerStatusResponseBody) *GetBatchChangeTableOwnerStatusResponse {
	s.Body = v
	return s
}

func (s *GetBatchChangeTableOwnerStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
