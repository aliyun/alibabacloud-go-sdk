// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse
	GetBody() *ModifyDBClusterDescriptionResponseBody
}

type ModifyDBClusterDescriptionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterDescriptionResponse) GetBody() *ModifyDBClusterDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDBClusterDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetStatusCode(v int32) *ModifyDBClusterDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
