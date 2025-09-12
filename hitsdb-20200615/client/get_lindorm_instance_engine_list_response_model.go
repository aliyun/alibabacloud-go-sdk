// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceEngineListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormInstanceEngineListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormInstanceEngineListResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormInstanceEngineListResponseBody) *GetLindormInstanceEngineListResponse
	GetBody() *GetLindormInstanceEngineListResponseBody
}

type GetLindormInstanceEngineListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormInstanceEngineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormInstanceEngineListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceEngineListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormInstanceEngineListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormInstanceEngineListResponse) GetBody() *GetLindormInstanceEngineListResponseBody {
	return s.Body
}

func (s *GetLindormInstanceEngineListResponse) SetHeaders(v map[string]*string) *GetLindormInstanceEngineListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceEngineListResponse) SetStatusCode(v int32) *GetLindormInstanceEngineListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormInstanceEngineListResponse) SetBody(v *GetLindormInstanceEngineListResponseBody) *GetLindormInstanceEngineListResponse {
	s.Body = v
	return s
}

func (s *GetLindormInstanceEngineListResponse) Validate() error {
	return dara.Validate(s)
}
