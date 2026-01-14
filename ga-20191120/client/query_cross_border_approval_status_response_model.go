// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCrossBorderApprovalStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCrossBorderApprovalStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCrossBorderApprovalStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryCrossBorderApprovalStatusResponseBody) *QueryCrossBorderApprovalStatusResponse
	GetBody() *QueryCrossBorderApprovalStatusResponseBody
}

type QueryCrossBorderApprovalStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCrossBorderApprovalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCrossBorderApprovalStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCrossBorderApprovalStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryCrossBorderApprovalStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCrossBorderApprovalStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCrossBorderApprovalStatusResponse) GetBody() *QueryCrossBorderApprovalStatusResponseBody {
	return s.Body
}

func (s *QueryCrossBorderApprovalStatusResponse) SetHeaders(v map[string]*string) *QueryCrossBorderApprovalStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryCrossBorderApprovalStatusResponse) SetStatusCode(v int32) *QueryCrossBorderApprovalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCrossBorderApprovalStatusResponse) SetBody(v *QueryCrossBorderApprovalStatusResponseBody) *QueryCrossBorderApprovalStatusResponse {
	s.Body = v
	return s
}

func (s *QueryCrossBorderApprovalStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
