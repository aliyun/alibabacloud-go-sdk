// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagVbrRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagVbrRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagVbrRelationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagVbrRelationsResponseBody) *DescribeSagVbrRelationsResponse
	GetBody() *DescribeSagVbrRelationsResponseBody
}

type DescribeSagVbrRelationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagVbrRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagVbrRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagVbrRelationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagVbrRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagVbrRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagVbrRelationsResponse) GetBody() *DescribeSagVbrRelationsResponseBody {
	return s.Body
}

func (s *DescribeSagVbrRelationsResponse) SetHeaders(v map[string]*string) *DescribeSagVbrRelationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagVbrRelationsResponse) SetStatusCode(v int32) *DescribeSagVbrRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagVbrRelationsResponse) SetBody(v *DescribeSagVbrRelationsResponseBody) *DescribeSagVbrRelationsResponse {
	s.Body = v
	return s
}

func (s *DescribeSagVbrRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
