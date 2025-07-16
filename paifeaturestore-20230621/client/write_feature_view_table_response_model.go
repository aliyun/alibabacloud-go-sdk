// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWriteFeatureViewTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WriteFeatureViewTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WriteFeatureViewTableResponse
	GetStatusCode() *int32
	SetBody(v *WriteFeatureViewTableResponseBody) *WriteFeatureViewTableResponse
	GetBody() *WriteFeatureViewTableResponseBody
}

type WriteFeatureViewTableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WriteFeatureViewTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WriteFeatureViewTableResponse) String() string {
	return dara.Prettify(s)
}

func (s WriteFeatureViewTableResponse) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WriteFeatureViewTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WriteFeatureViewTableResponse) GetBody() *WriteFeatureViewTableResponseBody {
	return s.Body
}

func (s *WriteFeatureViewTableResponse) SetHeaders(v map[string]*string) *WriteFeatureViewTableResponse {
	s.Headers = v
	return s
}

func (s *WriteFeatureViewTableResponse) SetStatusCode(v int32) *WriteFeatureViewTableResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteFeatureViewTableResponse) SetBody(v *WriteFeatureViewTableResponseBody) *WriteFeatureViewTableResponse {
	s.Body = v
	return s
}

func (s *WriteFeatureViewTableResponse) Validate() error {
	return dara.Validate(s)
}
