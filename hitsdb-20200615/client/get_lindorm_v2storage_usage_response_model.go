// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2StorageUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2StorageUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2StorageUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2StorageUsageResponseBody) *GetLindormV2StorageUsageResponse
	GetBody() *GetLindormV2StorageUsageResponseBody
}

type GetLindormV2StorageUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2StorageUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2StorageUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StorageUsageResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2StorageUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2StorageUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2StorageUsageResponse) GetBody() *GetLindormV2StorageUsageResponseBody {
	return s.Body
}

func (s *GetLindormV2StorageUsageResponse) SetHeaders(v map[string]*string) *GetLindormV2StorageUsageResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2StorageUsageResponse) SetStatusCode(v int32) *GetLindormV2StorageUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2StorageUsageResponse) SetBody(v *GetLindormV2StorageUsageResponseBody) *GetLindormV2StorageUsageResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2StorageUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
