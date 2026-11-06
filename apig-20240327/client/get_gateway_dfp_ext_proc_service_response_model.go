// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDfpExtProcServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayDfpExtProcServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayDfpExtProcServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayDfpExtProcServiceResponseBody) *GetGatewayDfpExtProcServiceResponse
	GetBody() *GetGatewayDfpExtProcServiceResponseBody
}

type GetGatewayDfpExtProcServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayDfpExtProcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayDfpExtProcServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDfpExtProcServiceResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayDfpExtProcServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayDfpExtProcServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayDfpExtProcServiceResponse) GetBody() *GetGatewayDfpExtProcServiceResponseBody {
	return s.Body
}

func (s *GetGatewayDfpExtProcServiceResponse) SetHeaders(v map[string]*string) *GetGatewayDfpExtProcServiceResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponse) SetStatusCode(v int32) *GetGatewayDfpExtProcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponse) SetBody(v *GetGatewayDfpExtProcServiceResponseBody) *GetGatewayDfpExtProcServiceResponse {
	s.Body = v
	return s
}

func (s *GetGatewayDfpExtProcServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
