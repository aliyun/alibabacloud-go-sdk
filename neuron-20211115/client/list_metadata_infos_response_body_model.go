// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetadataInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMetadataInfosResponseBody
	GetRequestId() *string
	SetResult(v *MetadataInfoListResult) *ListMetadataInfosResponseBody
	GetResult() *MetadataInfoListResult
}

type ListMetadataInfosResponseBody struct {
	RequestId *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *MetadataInfoListResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListMetadataInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetadataInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetadataInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetadataInfosResponseBody) GetResult() *MetadataInfoListResult {
	return s.Result
}

func (s *ListMetadataInfosResponseBody) SetRequestId(v string) *ListMetadataInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetadataInfosResponseBody) SetResult(v *MetadataInfoListResult) *ListMetadataInfosResponseBody {
	s.Result = v
	return s
}

func (s *ListMetadataInfosResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
