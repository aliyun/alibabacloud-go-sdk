// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMessageContactUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelMessageContactUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelMessageContactUpdateResponse
	GetStatusCode() *int32
	SetBody(v *CancelMessageContactUpdateResponseBody) *CancelMessageContactUpdateResponse
	GetBody() *CancelMessageContactUpdateResponseBody
}

type CancelMessageContactUpdateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelMessageContactUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelMessageContactUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelMessageContactUpdateResponse) GoString() string {
	return s.String()
}

func (s *CancelMessageContactUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelMessageContactUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelMessageContactUpdateResponse) GetBody() *CancelMessageContactUpdateResponseBody {
	return s.Body
}

func (s *CancelMessageContactUpdateResponse) SetHeaders(v map[string]*string) *CancelMessageContactUpdateResponse {
	s.Headers = v
	return s
}

func (s *CancelMessageContactUpdateResponse) SetStatusCode(v int32) *CancelMessageContactUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelMessageContactUpdateResponse) SetBody(v *CancelMessageContactUpdateResponseBody) *CancelMessageContactUpdateResponse {
	s.Body = v
	return s
}

func (s *CancelMessageContactUpdateResponse) Validate() error {
	return dara.Validate(s)
}
