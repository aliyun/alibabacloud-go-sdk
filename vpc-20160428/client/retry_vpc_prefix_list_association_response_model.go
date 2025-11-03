// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVpcPrefixListAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryVpcPrefixListAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryVpcPrefixListAssociationResponse
	GetStatusCode() *int32
	SetBody(v *RetryVpcPrefixListAssociationResponseBody) *RetryVpcPrefixListAssociationResponse
	GetBody() *RetryVpcPrefixListAssociationResponseBody
}

type RetryVpcPrefixListAssociationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryVpcPrefixListAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryVpcPrefixListAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryVpcPrefixListAssociationResponse) GoString() string {
	return s.String()
}

func (s *RetryVpcPrefixListAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryVpcPrefixListAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryVpcPrefixListAssociationResponse) GetBody() *RetryVpcPrefixListAssociationResponseBody {
	return s.Body
}

func (s *RetryVpcPrefixListAssociationResponse) SetHeaders(v map[string]*string) *RetryVpcPrefixListAssociationResponse {
	s.Headers = v
	return s
}

func (s *RetryVpcPrefixListAssociationResponse) SetStatusCode(v int32) *RetryVpcPrefixListAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryVpcPrefixListAssociationResponse) SetBody(v *RetryVpcPrefixListAssociationResponseBody) *RetryVpcPrefixListAssociationResponse {
	s.Body = v
	return s
}

func (s *RetryVpcPrefixListAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
