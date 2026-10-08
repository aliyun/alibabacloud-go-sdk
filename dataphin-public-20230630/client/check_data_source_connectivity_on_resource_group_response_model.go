// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityOnResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDataSourceConnectivityOnResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDataSourceConnectivityOnResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CheckDataSourceConnectivityOnResourceGroupResponseBody) *CheckDataSourceConnectivityOnResourceGroupResponse
	GetBody() *CheckDataSourceConnectivityOnResourceGroupResponseBody
}

type CheckDataSourceConnectivityOnResourceGroupResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataSourceConnectivityOnResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataSourceConnectivityOnResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityOnResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) GetBody() *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	return s.Body
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) SetHeaders(v map[string]*string) *CheckDataSourceConnectivityOnResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) SetStatusCode(v int32) *CheckDataSourceConnectivityOnResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) SetBody(v *CheckDataSourceConnectivityOnResourceGroupResponseBody) *CheckDataSourceConnectivityOnResourceGroupResponse {
	s.Body = v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
