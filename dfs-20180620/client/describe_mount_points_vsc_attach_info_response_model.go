// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountPointsVscAttachInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMountPointsVscAttachInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMountPointsVscAttachInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMountPointsVscAttachInfoResponseBody) *DescribeMountPointsVscAttachInfoResponse
	GetBody() *DescribeMountPointsVscAttachInfoResponseBody
}

type DescribeMountPointsVscAttachInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMountPointsVscAttachInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMountPointsVscAttachInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountPointsVscAttachInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountPointsVscAttachInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMountPointsVscAttachInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMountPointsVscAttachInfoResponse) GetBody() *DescribeMountPointsVscAttachInfoResponseBody {
	return s.Body
}

func (s *DescribeMountPointsVscAttachInfoResponse) SetHeaders(v map[string]*string) *DescribeMountPointsVscAttachInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponse) SetStatusCode(v int32) *DescribeMountPointsVscAttachInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponse) SetBody(v *DescribeMountPointsVscAttachInfoResponseBody) *DescribeMountPointsVscAttachInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
