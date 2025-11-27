// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossBackupMetaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossBackupMetaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossBackupMetaListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossBackupMetaListResponseBody) *DescribeCrossBackupMetaListResponse
	GetBody() *DescribeCrossBackupMetaListResponseBody
}

type DescribeCrossBackupMetaListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossBackupMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossBackupMetaListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossBackupMetaListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossBackupMetaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossBackupMetaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossBackupMetaListResponse) GetBody() *DescribeCrossBackupMetaListResponseBody {
	return s.Body
}

func (s *DescribeCrossBackupMetaListResponse) SetHeaders(v map[string]*string) *DescribeCrossBackupMetaListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossBackupMetaListResponse) SetStatusCode(v int32) *DescribeCrossBackupMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossBackupMetaListResponse) SetBody(v *DescribeCrossBackupMetaListResponseBody) *DescribeCrossBackupMetaListResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossBackupMetaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
