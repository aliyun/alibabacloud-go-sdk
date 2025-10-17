// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterDescriptionZonalResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterDescriptionZonalResponseBody) *ModifyDBClusterDescriptionZonalResponse
	GetBody() *ModifyDBClusterDescriptionZonalResponseBody
}

type ModifyDBClusterDescriptionZonalResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterDescriptionZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterDescriptionZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionZonalResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterDescriptionZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterDescriptionZonalResponse) GetBody() *ModifyDBClusterDescriptionZonalResponseBody {
	return s.Body
}

func (s *ModifyDBClusterDescriptionZonalResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionZonalResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDescriptionZonalResponse) SetStatusCode(v int32) *ModifyDBClusterDescriptionZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalResponse) SetBody(v *ModifyDBClusterDescriptionZonalResponseBody) *ModifyDBClusterDescriptionZonalResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterDescriptionZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
