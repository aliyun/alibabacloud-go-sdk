// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateContactGroupResponseBody) *CreateOrUpdateContactGroupResponse
	GetBody() *CreateOrUpdateContactGroupResponseBody
}

type CreateOrUpdateContactGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateContactGroupResponse) GetBody() *CreateOrUpdateContactGroupResponseBody {
	return s.Body
}

func (s *CreateOrUpdateContactGroupResponse) SetHeaders(v map[string]*string) *CreateOrUpdateContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateContactGroupResponse) SetStatusCode(v int32) *CreateOrUpdateContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponse) SetBody(v *CreateOrUpdateContactGroupResponseBody) *CreateOrUpdateContactGroupResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateContactGroupResponse) Validate() error {
	return dara.Validate(s)
}
