// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetSwitchInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDatasetSwitchInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDatasetSwitchInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryDatasetSwitchInfoResponseBody) *QueryDatasetSwitchInfoResponse
	GetBody() *QueryDatasetSwitchInfoResponseBody
}

type QueryDatasetSwitchInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetSwitchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDatasetSwitchInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetSwitchInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDatasetSwitchInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDatasetSwitchInfoResponse) GetBody() *QueryDatasetSwitchInfoResponseBody {
	return s.Body
}

func (s *QueryDatasetSwitchInfoResponse) SetHeaders(v map[string]*string) *QueryDatasetSwitchInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetSwitchInfoResponse) SetStatusCode(v int32) *QueryDatasetSwitchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponse) SetBody(v *QueryDatasetSwitchInfoResponseBody) *QueryDatasetSwitchInfoResponse {
	s.Body = v
	return s
}

func (s *QueryDatasetSwitchInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
