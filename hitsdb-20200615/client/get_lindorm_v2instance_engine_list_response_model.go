// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceEngineListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2InstanceEngineListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2InstanceEngineListResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2InstanceEngineListResponseBody) *GetLindormV2InstanceEngineListResponse
	GetBody() *GetLindormV2InstanceEngineListResponseBody
}

type GetLindormV2InstanceEngineListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceEngineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceEngineListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2InstanceEngineListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2InstanceEngineListResponse) GetBody() *GetLindormV2InstanceEngineListResponseBody {
	return s.Body
}

func (s *GetLindormV2InstanceEngineListResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceEngineListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceEngineListResponse) SetStatusCode(v int32) *GetLindormV2InstanceEngineListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponse) SetBody(v *GetLindormV2InstanceEngineListResponseBody) *GetLindormV2InstanceEngineListResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2InstanceEngineListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
