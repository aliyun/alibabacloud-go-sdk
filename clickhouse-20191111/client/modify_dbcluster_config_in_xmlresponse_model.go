// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigInXMLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterConfigInXMLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterConfigInXMLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterConfigInXMLResponseBody) *ModifyDBClusterConfigInXMLResponse
	GetBody() *ModifyDBClusterConfigInXMLResponseBody
}

type ModifyDBClusterConfigInXMLResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterConfigInXMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterConfigInXMLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigInXMLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigInXMLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterConfigInXMLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterConfigInXMLResponse) GetBody() *ModifyDBClusterConfigInXMLResponseBody {
	return s.Body
}

func (s *ModifyDBClusterConfigInXMLResponse) SetHeaders(v map[string]*string) *ModifyDBClusterConfigInXMLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterConfigInXMLResponse) SetStatusCode(v int32) *ModifyDBClusterConfigInXMLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLResponse) SetBody(v *ModifyDBClusterConfigInXMLResponseBody) *ModifyDBClusterConfigInXMLResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterConfigInXMLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
