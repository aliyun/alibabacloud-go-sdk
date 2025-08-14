// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveEditingIndexFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveEditingIndexFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveEditingIndexFileResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveEditingIndexFileResponseBody) *GetLiveEditingIndexFileResponse
	GetBody() *GetLiveEditingIndexFileResponseBody
}

type GetLiveEditingIndexFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveEditingIndexFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveEditingIndexFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingIndexFileResponse) GoString() string {
	return s.String()
}

func (s *GetLiveEditingIndexFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveEditingIndexFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveEditingIndexFileResponse) GetBody() *GetLiveEditingIndexFileResponseBody {
	return s.Body
}

func (s *GetLiveEditingIndexFileResponse) SetHeaders(v map[string]*string) *GetLiveEditingIndexFileResponse {
	s.Headers = v
	return s
}

func (s *GetLiveEditingIndexFileResponse) SetStatusCode(v int32) *GetLiveEditingIndexFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveEditingIndexFileResponse) SetBody(v *GetLiveEditingIndexFileResponseBody) *GetLiveEditingIndexFileResponse {
	s.Body = v
	return s
}

func (s *GetLiveEditingIndexFileResponse) Validate() error {
	return dara.Validate(s)
}
