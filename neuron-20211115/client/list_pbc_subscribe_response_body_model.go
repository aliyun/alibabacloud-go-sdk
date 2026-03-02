// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcSubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPbcListResult(v *PbcListResult) *ListPbcSubscribeResponseBody
	GetPbcListResult() *PbcListResult
	SetRequestId(v string) *ListPbcSubscribeResponseBody
	GetRequestId() *string
}

type ListPbcSubscribeResponseBody struct {
	PbcListResult *PbcListResult `json:"pbcListResult,omitempty" xml:"pbcListResult,omitempty"`
	RequestId     *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPbcSubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPbcSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *ListPbcSubscribeResponseBody) GetPbcListResult() *PbcListResult {
	return s.PbcListResult
}

func (s *ListPbcSubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPbcSubscribeResponseBody) SetPbcListResult(v *PbcListResult) *ListPbcSubscribeResponseBody {
	s.PbcListResult = v
	return s
}

func (s *ListPbcSubscribeResponseBody) SetRequestId(v string) *ListPbcSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPbcSubscribeResponseBody) Validate() error {
	if s.PbcListResult != nil {
		if err := s.PbcListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
