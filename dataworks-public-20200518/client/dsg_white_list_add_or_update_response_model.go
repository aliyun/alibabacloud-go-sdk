// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListAddOrUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgWhiteListAddOrUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgWhiteListAddOrUpdateResponse
	GetStatusCode() *int32
	SetBody(v *DsgWhiteListAddOrUpdateResponseBody) *DsgWhiteListAddOrUpdateResponse
	GetBody() *DsgWhiteListAddOrUpdateResponseBody
}

type DsgWhiteListAddOrUpdateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgWhiteListAddOrUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgWhiteListAddOrUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListAddOrUpdateResponse) GoString() string {
	return s.String()
}

func (s *DsgWhiteListAddOrUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgWhiteListAddOrUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgWhiteListAddOrUpdateResponse) GetBody() *DsgWhiteListAddOrUpdateResponseBody {
	return s.Body
}

func (s *DsgWhiteListAddOrUpdateResponse) SetHeaders(v map[string]*string) *DsgWhiteListAddOrUpdateResponse {
	s.Headers = v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponse) SetStatusCode(v int32) *DsgWhiteListAddOrUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponse) SetBody(v *DsgWhiteListAddOrUpdateResponseBody) *DsgWhiteListAddOrUpdateResponse {
	s.Body = v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponse) Validate() error {
	return dara.Validate(s)
}
