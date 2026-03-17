// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointNetworkQualitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessPointNetworkQualitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessPointNetworkQualitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessPointNetworkQualitiesResponseBody) *ListAccessPointNetworkQualitiesResponse
	GetBody() *ListAccessPointNetworkQualitiesResponseBody
}

type ListAccessPointNetworkQualitiesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessPointNetworkQualitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessPointNetworkQualitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointNetworkQualitiesResponse) GoString() string {
	return s.String()
}

func (s *ListAccessPointNetworkQualitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessPointNetworkQualitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessPointNetworkQualitiesResponse) GetBody() *ListAccessPointNetworkQualitiesResponseBody {
	return s.Body
}

func (s *ListAccessPointNetworkQualitiesResponse) SetHeaders(v map[string]*string) *ListAccessPointNetworkQualitiesResponse {
	s.Headers = v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponse) SetStatusCode(v int32) *ListAccessPointNetworkQualitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponse) SetBody(v *ListAccessPointNetworkQualitiesResponseBody) *ListAccessPointNetworkQualitiesResponse {
	s.Body = v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
