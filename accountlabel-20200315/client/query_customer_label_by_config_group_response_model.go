// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerLabelByConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomerLabelByConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomerLabelByConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomerLabelByConfigGroupResponseBody) *QueryCustomerLabelByConfigGroupResponse
	GetBody() *QueryCustomerLabelByConfigGroupResponseBody
}

type QueryCustomerLabelByConfigGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerLabelByConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerLabelByConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelByConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelByConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomerLabelByConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomerLabelByConfigGroupResponse) GetBody() *QueryCustomerLabelByConfigGroupResponseBody {
	return s.Body
}

func (s *QueryCustomerLabelByConfigGroupResponse) SetHeaders(v map[string]*string) *QueryCustomerLabelByConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponse) SetStatusCode(v int32) *QueryCustomerLabelByConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponse) SetBody(v *QueryCustomerLabelByConfigGroupResponseBody) *QueryCustomerLabelByConfigGroupResponse {
	s.Body = v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
