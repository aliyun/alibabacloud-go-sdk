// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2InstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2InstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2InstanceResponseBody) *GetLindormV2InstanceResponse
	GetBody() *GetLindormV2InstanceResponseBody
}

type GetLindormV2InstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2InstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2InstanceResponse) GetBody() *GetLindormV2InstanceResponseBody {
	return s.Body
}

func (s *GetLindormV2InstanceResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceResponse) SetStatusCode(v int32) *GetLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceResponse) SetBody(v *GetLindormV2InstanceResponseBody) *GetLindormV2InstanceResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2InstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
