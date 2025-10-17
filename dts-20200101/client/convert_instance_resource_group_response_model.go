// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertInstanceResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertInstanceResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ConvertInstanceResourceGroupResponseBody) *ConvertInstanceResourceGroupResponse
	GetBody() *ConvertInstanceResourceGroupResponseBody
}

type ConvertInstanceResourceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertInstanceResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertInstanceResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertInstanceResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertInstanceResourceGroupResponse) GetBody() *ConvertInstanceResourceGroupResponseBody {
	return s.Body
}

func (s *ConvertInstanceResourceGroupResponse) SetHeaders(v map[string]*string) *ConvertInstanceResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ConvertInstanceResourceGroupResponse) SetStatusCode(v int32) *ConvertInstanceResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertInstanceResourceGroupResponse) SetBody(v *ConvertInstanceResourceGroupResponseBody) *ConvertInstanceResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ConvertInstanceResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
