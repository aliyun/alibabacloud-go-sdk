// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2InstanceDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2InstanceDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2InstanceDetailsResponseBody) *GetLindormV2InstanceDetailsResponse
	GetBody() *GetLindormV2InstanceDetailsResponseBody
}

type GetLindormV2InstanceDetailsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2InstanceDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2InstanceDetailsResponse) GetBody() *GetLindormV2InstanceDetailsResponseBody {
	return s.Body
}

func (s *GetLindormV2InstanceDetailsResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponse) SetStatusCode(v int32) *GetLindormV2InstanceDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponse) SetBody(v *GetLindormV2InstanceDetailsResponseBody) *GetLindormV2InstanceDetailsResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
