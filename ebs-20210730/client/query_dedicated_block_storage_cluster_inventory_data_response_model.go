// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDedicatedBlockStorageClusterInventoryDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDedicatedBlockStorageClusterInventoryDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDedicatedBlockStorageClusterInventoryDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) *QueryDedicatedBlockStorageClusterInventoryDataResponse
	GetBody() *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
}

type QueryDedicatedBlockStorageClusterInventoryDataResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDedicatedBlockStorageClusterInventoryDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) GetBody() *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	return s.Body
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) SetHeaders(v map[string]*string) *QueryDedicatedBlockStorageClusterInventoryDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) SetStatusCode(v int32) *QueryDedicatedBlockStorageClusterInventoryDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) SetBody(v *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) *QueryDedicatedBlockStorageClusterInventoryDataResponse {
	s.Body = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
