// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsUnfreezeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOssCheckResultsUnfreezeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOssCheckResultsUnfreezeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOssCheckResultsUnfreezeResponseBody) *UpdateOssCheckResultsUnfreezeResponse
	GetBody() *UpdateOssCheckResultsUnfreezeResponseBody
}

type UpdateOssCheckResultsUnfreezeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssCheckResultsUnfreezeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssCheckResultsUnfreezeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsUnfreezeResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsUnfreezeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOssCheckResultsUnfreezeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOssCheckResultsUnfreezeResponse) GetBody() *UpdateOssCheckResultsUnfreezeResponseBody {
	return s.Body
}

func (s *UpdateOssCheckResultsUnfreezeResponse) SetHeaders(v map[string]*string) *UpdateOssCheckResultsUnfreezeResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponse) SetStatusCode(v int32) *UpdateOssCheckResultsUnfreezeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponse) SetBody(v *UpdateOssCheckResultsUnfreezeResponseBody) *UpdateOssCheckResultsUnfreezeResponse {
	s.Body = v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
