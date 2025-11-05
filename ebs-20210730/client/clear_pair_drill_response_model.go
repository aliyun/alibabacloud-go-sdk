// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearPairDrillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearPairDrillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearPairDrillResponse
	GetStatusCode() *int32
	SetBody(v *ClearPairDrillResponseBody) *ClearPairDrillResponse
	GetBody() *ClearPairDrillResponseBody
}

type ClearPairDrillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearPairDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearPairDrillResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearPairDrillResponse) GoString() string {
	return s.String()
}

func (s *ClearPairDrillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearPairDrillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearPairDrillResponse) GetBody() *ClearPairDrillResponseBody {
	return s.Body
}

func (s *ClearPairDrillResponse) SetHeaders(v map[string]*string) *ClearPairDrillResponse {
	s.Headers = v
	return s
}

func (s *ClearPairDrillResponse) SetStatusCode(v int32) *ClearPairDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearPairDrillResponse) SetBody(v *ClearPairDrillResponseBody) *ClearPairDrillResponse {
	s.Body = v
	return s
}

func (s *ClearPairDrillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
