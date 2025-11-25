// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelKyuubiSparkApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelKyuubiSparkApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelKyuubiSparkApplicationResponse
	GetStatusCode() *int32
	SetBody(v *CancelKyuubiSparkApplicationResponseBody) *CancelKyuubiSparkApplicationResponse
	GetBody() *CancelKyuubiSparkApplicationResponseBody
}

type CancelKyuubiSparkApplicationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelKyuubiSparkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelKyuubiSparkApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelKyuubiSparkApplicationResponse) GoString() string {
	return s.String()
}

func (s *CancelKyuubiSparkApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelKyuubiSparkApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelKyuubiSparkApplicationResponse) GetBody() *CancelKyuubiSparkApplicationResponseBody {
	return s.Body
}

func (s *CancelKyuubiSparkApplicationResponse) SetHeaders(v map[string]*string) *CancelKyuubiSparkApplicationResponse {
	s.Headers = v
	return s
}

func (s *CancelKyuubiSparkApplicationResponse) SetStatusCode(v int32) *CancelKyuubiSparkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelKyuubiSparkApplicationResponse) SetBody(v *CancelKyuubiSparkApplicationResponseBody) *CancelKyuubiSparkApplicationResponse {
	s.Body = v
	return s
}

func (s *CancelKyuubiSparkApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
