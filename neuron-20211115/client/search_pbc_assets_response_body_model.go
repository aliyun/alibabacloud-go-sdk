// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPbcAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchPbcAssetsResponseBody
	GetRequestId() *string
	SetResult(v *PbcListResult) *SearchPbcAssetsResponseBody
	GetResult() *PbcListResult
}

type SearchPbcAssetsResponseBody struct {
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *PbcListResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SearchPbcAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchPbcAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPbcAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchPbcAssetsResponseBody) GetResult() *PbcListResult {
	return s.Result
}

func (s *SearchPbcAssetsResponseBody) SetRequestId(v string) *SearchPbcAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPbcAssetsResponseBody) SetResult(v *PbcListResult) *SearchPbcAssetsResponseBody {
	s.Result = v
	return s
}

func (s *SearchPbcAssetsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
