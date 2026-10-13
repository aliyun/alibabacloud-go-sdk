// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCompareIncrementDbOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaCompareIncrementDbOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaCompareIncrementDbOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaCompareIncrementDbOverviewResponseBody) *GetMetaCompareIncrementDbOverviewResponse
	GetBody() *GetMetaCompareIncrementDbOverviewResponseBody
}

type GetMetaCompareIncrementDbOverviewResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaCompareIncrementDbOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaCompareIncrementDbOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCompareIncrementDbOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetMetaCompareIncrementDbOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaCompareIncrementDbOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaCompareIncrementDbOverviewResponse) GetBody() *GetMetaCompareIncrementDbOverviewResponseBody {
	return s.Body
}

func (s *GetMetaCompareIncrementDbOverviewResponse) SetHeaders(v map[string]*string) *GetMetaCompareIncrementDbOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponse) SetStatusCode(v int32) *GetMetaCompareIncrementDbOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponse) SetBody(v *GetMetaCompareIncrementDbOverviewResponseBody) *GetMetaCompareIncrementDbOverviewResponse {
	s.Body = v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
