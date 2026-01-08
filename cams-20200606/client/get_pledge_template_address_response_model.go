// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPledgeTemplateAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPledgeTemplateAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPledgeTemplateAddressResponse
	GetStatusCode() *int32
	SetBody(v *GetPledgeTemplateAddressResponseBody) *GetPledgeTemplateAddressResponse
	GetBody() *GetPledgeTemplateAddressResponseBody
}

type GetPledgeTemplateAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPledgeTemplateAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPledgeTemplateAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPledgeTemplateAddressResponse) GoString() string {
	return s.String()
}

func (s *GetPledgeTemplateAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPledgeTemplateAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPledgeTemplateAddressResponse) GetBody() *GetPledgeTemplateAddressResponseBody {
	return s.Body
}

func (s *GetPledgeTemplateAddressResponse) SetHeaders(v map[string]*string) *GetPledgeTemplateAddressResponse {
	s.Headers = v
	return s
}

func (s *GetPledgeTemplateAddressResponse) SetStatusCode(v int32) *GetPledgeTemplateAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPledgeTemplateAddressResponse) SetBody(v *GetPledgeTemplateAddressResponseBody) *GetPledgeTemplateAddressResponse {
	s.Body = v
	return s
}

func (s *GetPledgeTemplateAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
